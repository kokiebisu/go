package example3

type Person struct {
	Name string
	Age  int
}

// Displays Age
func DisplayAge(person Person) int {
	return person.Age
}
