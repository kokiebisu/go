package example4

// A structure of a person
type Person struct {
	Name string
	Age  int
}

// A structure of a team
type Team struct {
	Name   string
	People []Person
}

// Retrieves a team
func (t Team) GetTeam() Team {
	people := []Person{Person{Name: "Ken", Age: 21}, Person{Name: "Ben", Age: 18}}
	team := Team{Name: "the best team", People: people}
	return team
}

// Creates a team
func CreateTeam(name string, people []Person) Team {
	team := Team{Name: name, People: people}
	return team
}
