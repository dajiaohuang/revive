package fixtures

var duplicateSwitchCondition = func(a bool) int {
	switch {
	case a:
		return 1
	case a: // MATCH /case clause at line 5 has the same condition/
		return 2
	}
	return 0
}

func namedDuplicateSwitchCondition(a bool) int {
	switch {
	case a:
		return 1
	case a: // MATCH /case clause at line 15 has the same condition/
		return 2
	}
	return 0
}
