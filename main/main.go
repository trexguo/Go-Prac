package main

import "fmt"

func main() {
	numArray := []int {5, 4, 3, 2, 1}
	numArray1 := make([]int, 1, 5)

	fmt.Println(numArray1)
	copy(numArray1, numArray)
	fmt.Println(numArray1)
}