package main

import "fmt"

func main() {
	numArray := []int{5, 4, 3, 2, 1}
	numArray1 := make([]int, 5, 10)  // https://sanyuesha.com/2017/07/26/go-make-and-new/

	fmt.Println(numArray1)
	copy(numArray1, numArray)
	fmt.Println(numArray1)

	numArray1 = append(numArray1, 0, -1)

	fmt.Println(numArray1)

}
