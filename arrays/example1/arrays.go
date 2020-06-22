package example1

import "fmt"

type Arrays struct{}

func (a Arrays) Run() {

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	fmt.Println(days[0:3])
}
