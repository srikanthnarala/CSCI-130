package main

import "fmt"

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	greatest := max(9,43,56,89,147,238,891,367,739)
	fmt.Println(greatest)
}
