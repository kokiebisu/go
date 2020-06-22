package main

import "github.com/kokiebisu/go/methods/example1"

func main() {
	var employee example1.Employee
	employee.Name = "Ken"
	employee.UpdateName("Ben")
	employee.PrintName()
}
