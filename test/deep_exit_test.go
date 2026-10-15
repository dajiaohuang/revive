package test_test

import (
	"testing"

	"github.com/mgechev/revive/rule"
)

func TestDeepExitAliases(t *testing.T) {
	testRule(t, "deep_exit_aliases", &rule.DeepExitRule{})
}

func TestDeepExit(t *testing.T) {
	testRule(t, "deep_exit", &rule.DeepExitRule{})
	testRule(t, "deep_exit_test", &rule.DeepExitRule{})
}
