package fixtures

func localPrint() {
	print := func(string) {}
	print("message")
}

func parameterPrint(println func(string)) {
	println("message")
}

func builtinPrint() {
	print("message")   // MATCH /avoid using built-in function "print", replace it by "fmt.Fprint(os.Stderr, "message")"/
	println("message") // MATCH /avoid using built-in function "println", replace it by "fmt.Fprintln(os.Stderr, "message")"/
}
