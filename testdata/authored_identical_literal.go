package fixtures

var sink = func() {}
var compare = func(condition bool) {
	if condition { // MATCH /both branches of the if are identical/
		sink()
	} else {
		sink()
	}
}

func named(condition bool) {
	if condition { // MATCH /both branches of the if are identical/
		sink()
	} else {
		sink()
	}
}
