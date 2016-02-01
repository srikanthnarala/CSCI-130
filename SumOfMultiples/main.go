package main

import "fmt"

func main() {
	var total = 0

	for i := 0; i <1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Scan(i)
			total = total + i

		}

	}
	fmt.Println(total)
}
