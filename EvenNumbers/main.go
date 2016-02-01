package main

import "fmt"

func main() {
	fmt.Println("List of even numbers from 1 to 100")

	for numbers := 0; numbers <= 100; numbers++ {

		if numbers%2 == 0 {

			fmt.Print("\n",numbers)
		}
	}

}
