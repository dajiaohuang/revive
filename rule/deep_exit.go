package rule

import (
	"fmt"
	"go/ast"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/mgechev/revive/internal/astutils"
	"github.com/mgechev/revive/lint"
)

// DeepExitRule lints program exit in functions other than main or init.
type DeepExitRule struct{}

// Apply applies the rule to given file.
func (*DeepExitRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	var failures []lint.Failure
	onFailure := func(failure lint.Failure) {
		failures = append(failures, failure)
	}

	imports := map[string]string{}
	for _, spec := range file.AST.Imports {
		// The Go parser has already validated this string literal.
		path, _ := strconv.Unquote(spec.Path.Value)
		switch path {
		case "os", "syscall", "log", "flag":
			name := path
			if spec.Name != nil {
				name = spec.Name.Name
			}
			imports[name] = path
		}
	}
	w := &lintDeepExit{onFailure: onFailure, isTestFile: file.IsTest(), imports: imports}
	ast.Walk(w, file.AST)
	return failures
}

// Name returns the rule name.
func (*DeepExitRule) Name() string {
	return "deep-exit"
}

type lintDeepExit struct {
	onFailure  func(lint.Failure)
	isTestFile bool
	imports    map[string]string
}

func (w *lintDeepExit) Visit(node ast.Node) ast.Visitor {
	if fd, ok := node.(*ast.FuncDecl); ok {
		if w.mustIgnore(fd) {
			return nil // skip analysis of this function
		}

		return w
	}

	se, ok := node.(*ast.ExprStmt)
	if !ok {
		return w
	}
	ce, ok := se.X.(*ast.CallExpr)
	if !ok {
		return w
	}

	fc, ok := ce.Fun.(*ast.SelectorExpr)
	if !ok {
		return w
	}
	id, ok := fc.X.(*ast.Ident)
	if !ok {
		return w
	}

	pkg := w.importedPackage(id)
	fn := fc.Sel.Name
	var exits bool
	if pkg == "flag" && fn == "NewFlagSet" {
		exits = len(ce.Args) == 2 && w.isFlagExitOnError(ce.Args[1])
	} else {
		exits = astutils.IsCallToExitFunction(pkg, fn, ce.Args)
	}
	if exits {
		msg := fmt.Sprintf("calls to %s.%s only in main() or init() functions", pkg, fn)

		switch {
		case pkg == "flag" && fn == "Parse":
			msg += "; move the call or refactor to use flag.NewFlagSet with flag.ContinueOnError"
		case pkg == "flag" && fn == "NewFlagSet":
			msg = "calls to flag.NewFlagSet with flag.ExitOnError only in main() or init() functions"
		}
		w.onFailure(lint.Failure{
			Confidence: 1,
			Node:       ce,
			Category:   lint.FailureCategoryBadPractice,
			Failure:    msg,
		})
	}

	return w
}

func (w *lintDeepExit) importedPackage(id *ast.Ident) string {
	if id.Obj != nil && id.Obj.Kind != ast.Pkg {
		return "" // This identifier is a local binding, not the imported package.
	}
	return w.imports[id.Name]
}

func (w *lintDeepExit) isFlagExitOnError(expr ast.Expr) bool {
	selector, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	id, ok := selector.X.(*ast.Ident)
	return ok && w.importedPackage(id) == "flag" && selector.Sel.Name == "ExitOnError"
}

func (w *lintDeepExit) mustIgnore(fd *ast.FuncDecl) bool {
	fn := fd.Name.Name

	return fn == "init" || fn == "main" || w.isTestMain(fd) || w.isTestExample(fd)
}

func (w *lintDeepExit) isTestMain(fd *ast.FuncDecl) bool {
	return w.isTestFile && fd.Name.Name == "TestMain"
}

// isTestExample returns true if the function is a testable example function.
// See https://go.dev/blog/examples#examples-are-tests for more information.
//
// Inspired by https://github.com/golang/go/blob/go1.23.0/src/go/doc/example.go#L72-L77
func (w *lintDeepExit) isTestExample(fd *ast.FuncDecl) bool {
	if !w.isTestFile {
		return false
	}
	name := fd.Name.Name
	const prefix = "Example"
	if !strings.HasPrefix(name, prefix) {
		return false
	}
	if len(name) == len(prefix) { // "Example" is a package level example
		return len(fd.Type.Params.List) == 0
	}
	r, _ := utf8.DecodeRuneInString(name[len(prefix):])
	if unicode.IsLower(r) {
		return false
	}
	return len(fd.Type.Params.List) == 0
}
