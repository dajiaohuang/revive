package fixtures

import (
	f "flag"
	l "log"
	o "os"
	s "syscall"
)

func aliasExit() {
	o.Exit(1) // MATCH /calls to os.Exit only in main() or init() functions/
}

func aliasSyscallExit() {
	s.Exit(1) // MATCH /calls to syscall.Exit only in main() or init() functions/
}

func aliasFatal() {
	l.Fatal("failure") // MATCH /calls to log.Fatal only in main() or init() functions/
}

func aliasFlagParse() {
	f.Parse() // MATCH /calls to flag.Parse only in main() or init() functions; move the call or refactor to use flag.NewFlagSet with flag.ContinueOnError/
}

func aliasExitFlagSet() {
	f.NewFlagSet("command", f.ExitOnError) // MATCH /calls to flag.NewFlagSet with flag.ExitOnError only in main() or init() functions/
}

func aliasContinueFlagSet() {
	f.NewFlagSet("command", f.ContinueOnError)
}

func shadowedAlias() {
	o := struct{ Exit func(int) }{Exit: func(int) {}}
	o.Exit(1)
}

func main() {
	o.Exit(0)
}
