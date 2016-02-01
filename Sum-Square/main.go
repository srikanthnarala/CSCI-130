package main

import "fmt"	//PROJECT EULER PROBLEM
		//FINDING DIFFERENCE BETWEEN (SUM OF SQUARES FROM 1 TO 10) AND (SQUARE OF SUM OF 1 TO 10) NUMBERS
func main() {

	var total = 0	//total is sum of squares from 1 to 10 numbers
	var total1=0	//total1 is square of sum of 1 to 10 numbers
	var x int

	for i := 1; i <= 10; i++ {
		fmt.Scan(i)
		total = total + (i * i)
	}
	fmt.Println("sum of squares from 1 to 10 is", total)
	total1=0
	for i:=1;i<=10 ;i++{
		fmt.Scan(i)
		total1=total1+i
	}
	total1=(total1*total1)
	x=total1-total
	fmt.Print("square of sum of 10 numbers is",total1)
	fmt.Println("\n difference of square of sum and sum of squares from 1 to 10 numbers is",x)


}
