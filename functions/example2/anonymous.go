package example2

func AddOne() func() int {
	var number int

	return func() int {
		number++
		return number + 1
	}
}
