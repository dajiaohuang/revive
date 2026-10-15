package fixtures

var sink = func() {}
var compare = func(a, b bool) {
	if a { // MATCH /"if...else if" chain with identical branches (lines 5 and 7)/
		sink()
	} else if b {
		sink()
	}
}

func named(a, b bool) {
	if a { // MATCH /"if...else if" chain with identical branches (lines 13 and 15)/
		sink()
	} else if b {
		sink()
	}
}
