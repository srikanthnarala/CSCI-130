package main

import (
	"fmt"
	"os"
	"text/template"
)

type student struct {
	Name        string
	Id          int
	GPA         float64
	Major       string
	Scholorship  bool
}

func main() {
	p1 := student{
		Name:  "Srikanth",
		Id:    109545754,
		Major: "computer Science",
		GPA:   3.8}
	if p1.GPA > 3.5 {
		p1.Scholorship = true
	}
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	err = tpl.Execute(os.Stdout, p1)
	if err != nil {
		fmt.Println(err)
	}
}
