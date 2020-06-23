package main

import (
	"fmt"

	"github.com/kokiebisu/go/interfaces/example1"
	"github.com/kokiebisu/go/interfaces/example2"
	"github.com/kokiebisu/go/interfaces/example3"
)

func main() {
	fmt.Println("runs example1")
	example1.Run()

	fmt.Println("runs example2")
	var basePlayer example2.BaseGuitarist
	var acousticPlayer example2.AcousticGuitarist

	basePlayer.Name = "Ken"
	basePlayer.PlayGuitar()

	acousticPlayer.Name = "Ben"
	acousticPlayer.PlayGuitar()

	var guitarists []example2.Guitarist
	guitarists = append(guitarists, basePlayer)
	guitarists = append(guitarists, acousticPlayer)

	fmt.Println(guitarists)

	fmt.Println("example3")
	var employees = example3.GatherEmployees()
	fmt.Println(employees)
}
