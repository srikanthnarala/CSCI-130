package main

import "fmt"

func main() {
	for numbers := 0; numbers <= 100; numbers++ {
		if numbers%3 == 0 {
			fmt.Print("\n  FIZZ", numbers)
		} else if numbers%5 == 0 {
			fmt.Print("\n  BUZZ", numbers)
		} else if numbers%15 == 0 {
			fmt.Print("\n  FIZZBUZZ", numbers)
		} else {
			fmt.Println(numbers)
		}
	}
}
