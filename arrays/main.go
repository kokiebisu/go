package main

import (
	"fmt"

	"github.com/kokiebisu/go/arrays/example1"
	"github.com/kokiebisu/go/arrays/example2"
	"github.com/kokiebisu/go/arrays/example3"
	"github.com/kokiebisu/go/arrays/example4"
)

func main() {
	fmt.Println("Running test 1")
	var test1 example1.Arrays
	test1.Run()

	fmt.Println("Running test 2")
	var test2 example2.Map
	characteristics := test2.Run()
	fmt.Println(characteristics)

	fmt.Println("Running test 3")
	var test3 example3.Person
	test3.Age = 21
	test3.Name = "Ben"
	example3.DisplayAge(test3)

	var test4 example4.Team
	team := test4.GetTeam()
	fmt.Println(team.Name)

	name := "The new team"
	people := []example4.Person{example4.Person{Name: "new ken", Age: 25}, example4.Person{Name: "new Ben", Age: 45}}
	newTeam := example4.CreateTeam(name, people)

	fmt.Println(newTeam)

}
