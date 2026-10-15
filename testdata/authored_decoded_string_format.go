package fixtures

func expectA(string) {}

func literalStrings() {
	expectA("a")
	expectA("\x61")
	expectA("\u0061")
	expectA("\141")
	expectA(`a`)
	expectA("b")    // MATCH /must contain a/
	expectA("\x62") // MATCH /must contain a/
}
