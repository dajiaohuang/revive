package fixtures

var duplicateCondition = func(a bool) int {
	if a {
		return 1
	} else if a { // MATCH /"if...else if" chain with identical conditions (lines 4 and 6)/
		return 2
	}
	return 0
}

func namedDuplicateCondition(a bool) int {
	if a {
		return 1
	} else if a { // MATCH /"if...else if" chain with identical conditions (lines 13 and 15)/
		return 2
	}
	return 0
}
