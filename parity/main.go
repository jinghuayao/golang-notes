package main

import "fmt"

func main() {
	slice := newSlice()
	fmt.Println(slice)
	for _, item := range slice {
		if item % 2 == 0 {
			fmt.Printf("even number %v", item)
			fmt.Println()
		} else {
			fmt.Printf("odd number %v", item)
			fmt.Println()
		}
	}
}

func newSlice() []int {
	res := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return res
}