package main

import "fmt"

func dailyTemperatures(T []int) []int {
	type Tuple struct {
		Index int
		Temp  int
	}
	// creates a slice of integers of length T with zero value
	result := make([]int, len(T))
	// creates a stack
	stack := []Tuple{}

	for i := 0; i < len(T); i++ {
		// if there is a tuple in the struct && if the top tuple in the
		// stack have a temperature lower than the temperature of the index
		for (len(stack) > 0) && stack[len(stack)-1].Temp < T[i] {
			result[stack[len(stack)-1].Index] = i - stack[len(stack)-1].Index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, Tuple{Index: i, Temp: T[i]})
	}
	return result
}

func main() {
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temp))
}
