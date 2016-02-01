package main

import "fmt"

func main() {
	var largenumber int
	var smallnumber int
	fmt.Print("Enter large number:")
	fmt.Scan(&largenumber)
	fmt.Print("Enter Small number:")
	fmt.Scan(&smallnumber)
	fmt.Println("large", largenumber, "small", smallnumber)
	Reminder:=largenumber%smallnumber
	fmt.Println("Reminder is",Reminder)
}