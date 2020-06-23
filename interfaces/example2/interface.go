package example2

import "fmt"

type Guitarist interface {
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the guitar\n", b.Name)
}

func (a AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the guitar\n", a.Name)
}
