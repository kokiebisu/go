package example3

type Employee interface {
	GetLanguage() string
	GetAge() string
}

type Engineer struct {
	Age      int
	Name     string
	Language string
}

func (e Engineer) GetLanguage() string {
	return e.Name + "programs in" + e.Language
}

func (e Engineer) GetAge() string {
	return e.Name + "is" + string(e.Age) + "years old"
}

func GatherEmployees() []Employee {
	var employees []Employee
	elliot := Engineer{Name: "Elliot", Language: "GO", Age: 15}

	employees = append(employees, elliot)
	return employees
}
