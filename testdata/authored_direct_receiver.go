package fixtures

type Counter int

func (n Counter) increment() {
	n++ // MATCH /suspicious assignment to a by-value method receiver/
}

func (n Counter) decrement() {
	n-- // MATCH /suspicious assignment to a by-value method receiver/
}

func (n Counter) next() Counter {
	n++
	return n
}

func (n *Counter) mutate() {
	(*n)++
}

func (n Counter) shadowedParameter() {
	func(n int) { n++ }(0)
}
