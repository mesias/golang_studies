package main

import "fmt"

type pearson struct {
	firstName string
	lastName  string
}

func main() {
	alex := pearson{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
