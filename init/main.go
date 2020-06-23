package main

import "fmt"

// This gets called first
var WhatIsThe = AnswerToLife()

// WhatIsThe is initialized to 42
func AnswerToLife() int {
	return 42
}

// Init then gets called which reassigns 0
func init() {
	WhatIsThe = 0
}

func main() {
	if WhatIsThe == 0 {
		fmt.Println("It's all a lie.")
	}
}
