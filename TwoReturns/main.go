package main

import "fmt"

func halfneven(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func main() {
	half ,even := halfneven(53)
	fmt.Println(half, even)
}
