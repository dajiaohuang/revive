package test_test

import (
	"testing"

	"github.com/mgechev/revive/lint"
	"github.com/mgechev/revive/rule"
)

func TestAuthoredCommentBlankOverlap(t *testing.T) {
	testRule(t, "authored_comment_blank_overlap", &rule.FileLengthLimitRule{}, &lint.RuleConfig{
		Arguments: lint.Arguments{map[string]any{"max": int64(3), "skipComments": true, "skipBlankLines": true}},
	})
}

func TestAuthoredLexicalPrint(t *testing.T) {
	testRule(t, "authored_lexical_print", &rule.UseFmtPrintRule{})
}

func TestAuthoredFunctionLiterals(t *testing.T) {
	for _, tc := range []struct {
		name string
		rule lint.Rule
	}{
		{"authored_identical_literal", &rule.IdenticalBranchesRule{}},
		{"authored_identical_chain_literal", &rule.IdenticalIfElseIfBranchesRule{}},
		{"authored_identical_condition_literal", &rule.IdenticalIfElseIfConditionsRule{}},
		{"authored_identical_switch_literal", &rule.IdenticalSwitchBranchesRule{}},
		{"authored_identical_switch_condition_literal", &rule.IdenticalSwitchConditionsRule{}},
	} {
		t.Run(tc.name, func(t *testing.T) { testRule(t, tc.name, tc.rule) })
	}
}

func TestAuthoredPhysicalCommentDensity(t *testing.T) {
	testRule(t, "authored_physical_comment_density", &rule.CommentsDensityRule{}, &lint.RuleConfig{
		Arguments: lint.Arguments{int64(20)},
	})
}

func TestAuthoredPhysicalCommentTokens(t *testing.T) {
	testRule(t, "authored_physical_comment_tokens", &rule.CommentsDensityRule{}, &lint.RuleConfig{
		Arguments: lint.Arguments{int64(80)},
	})
}

func TestAuthoredGroupedBuiltinTypes(t *testing.T) {
	testRule(t, "authored_grouped_builtin_types", &rule.RedefinesBuiltinIDRule{})
}

func TestAuthoredDirectReceiverIncDec(t *testing.T) {
	testRule(t, "authored_direct_receiver", &rule.ModifiesValRecRule{})
}

func TestAuthoredDecodedStringFormat(t *testing.T) {
	testRule(t, "authored_decoded_string_format", &rule.StringFormatRule{}, &lint.RuleConfig{
		Arguments: lint.Arguments{[]any{"expectA", "/^a$/", "must contain a"}},
	})
}
