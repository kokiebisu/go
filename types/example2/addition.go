package example2

// Runs the Run method from the addition struct
type Addition struct {
}

// Adds some numbers
func (a Addition) Run() int {
	var men uint8
	men = 8

	var women uint8
	women = 4

	var people int

	// This will give an error
	// people = men + women

	people = int(men) + int(women)

	return people
}
