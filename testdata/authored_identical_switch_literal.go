package fixtures

var duplicateSwitch = func(a int) int {
	switch a { // MATCH /"switch" with identical branches (lines 5 and 7)/
	case 1:
		return 10
	case 2:
		return 10
	}
	return 0
}

func namedDuplicateSwitch(a int) int {
	switch a { // MATCH /"switch" with identical branches (lines 15 and 17)/
	case 1:
		return 10
	case 2:
		return 10
	}
	return 0
}
