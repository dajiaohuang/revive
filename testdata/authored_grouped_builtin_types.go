package fixtures

type (
	Counter int32
	int     int32  // MATCH /redefinition of the built-in type int/
	string  []byte // MATCH /redefinition of the built-in type string/
)
