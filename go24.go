package main

import "fmt"

func sum(nums ...int) int {
	fmt.Print(nums, " ")
	total_sum := 0
	for _, num := range nums {
		total_sum += num
	}
	return total_sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
}
