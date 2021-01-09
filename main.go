package main

import "fmt"

//functions
func main() {
	tums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 7, 6, 5, 4, 3, 2}
	sums := make([]int, 10)

	copy(sums, tums)

	fmt.Println(sums)
}
