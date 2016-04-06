package main

import (
	"fmt"
"math")

func swap (x,y string) (string,string){
return y,x
}

func main() {
	fmt.Println("hello world")
	fmt.Println("squareroot of 88 is", math.Sqrt(88))
	a,b :=swap("hello","world")
	fmt.Println(a,b)

}
