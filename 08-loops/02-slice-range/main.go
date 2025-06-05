package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
