package main

import (
	"fmt"
)

func halfneven(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}

func main() {
	half, even := halfneven(29)
	fmt.Println(half, even)
}
