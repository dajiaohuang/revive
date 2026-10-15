package rule

import (
	"fmt"
	"go/scanner"
	"go/token"
	"strings"

	"github.com/mgechev/revive/lint"
)

// CommentsDensityRule enforces a minimum comment / code relation.
type CommentsDensityRule struct {
	minimumCommentsDensity int64
}

const defaultMinimumCommentsPercentage = 0

var _ lint.ConfigurableRule = (*CommentsDensityRule)(nil)

// Configure validates the rule configuration, and configures the rule accordingly.
//
// Configuration implements the [lint.ConfigurableRule] interface.
func (r *CommentsDensityRule) Configure(arguments lint.Arguments) error {
	if len(arguments) < 1 {
		r.minimumCommentsDensity = defaultMinimumCommentsPercentage
		return nil
	}

	var ok bool
	r.minimumCommentsDensity, ok = arguments[0].(int64)
	if !ok {
		return fmt.Errorf("invalid argument for %q rule: argument should be an int, got %T", r.Name(), arguments[0])
	}
	return nil
}

// Apply applies the rule to given file.
func (r *CommentsDensityRule) Apply(file *lint.File, _ lint.Arguments) []lint.Failure {
	commentsLines, codeLines := countSourceLines(file)
	density := (float32(commentsLines) / float32(codeLines+commentsLines)) * 100

	if density < float32(r.minimumCommentsDensity) {
		return []lint.Failure{
			{
				Category:   lint.FailureCategoryMaintenance,
				Node:       file.AST,
				Confidence: 1,
				Failure: fmt.Sprintf("the file has a comment density of %2.f%% (%d comment lines for %d code lines) but expected a minimum of %d%%",
					density, commentsLines, codeLines, r.minimumCommentsDensity),
			},
		}
	}

	return nil
}

// Name returns the rule name.
func (*CommentsDensityRule) Name() string {
	return "comments-density"
}

// countSourceLines counts physical lines occupied by comments and by Go tokens.
func countSourceLines(file *lint.File) (comments, code int) {
	source := file.Content()
	tokenFile := token.NewFileSet().AddFile(file.Name, -1, len(source))
	var lexer scanner.Scanner
	lexer.Init(tokenFile, source, nil, scanner.ScanComments)
	codeLines := map[int]struct{}{}
	commentLines := map[int]struct{}{}
	for {
		pos, kind, literal := lexer.Scan()
		if kind == token.EOF {
			break
		}
		if kind == token.SEMICOLON && literal == "\n" {
			continue // Inserted semicolons do not occupy an additional source line.
		}
		lines := codeLines
		if kind == token.COMMENT {
			lines = commentLines
		}
		start := tokenFile.PositionFor(pos, false).Line
		end := start + strings.Count(literal, "\n")
		for line := start; line <= end; line++ {
			lines[line] = struct{}{}
		}
	}
	return len(commentLines), len(codeLines)
}
