package main

import "fmt"

func fourSumCount(A []int, B []int, C []int, D []int) int {
	// first i would get all the possible combinations of the indexes
	// map{}
	mem := make(map[int]int)
	for _, a := range A {
		for _, b := range B {
			// increase the hash map value if more than 1 pair sums to the same value
			mem[a+b]++
		}
	}

	var output, sum int
	for _, c := range C {
		for _, d := range D {
			sum = (c + d)
			if v, ok := mem[sum]; ok {
				output += v
			}
		}
	}

	return output
}

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}
