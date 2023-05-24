package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// var (
	// 	name    = "batuhan"
	// 	surname = "yigit"
	// )

	var person Person
	person.Name = "batuhan"
	person.Age = 123

	fmt.Println("Hello ", person.Name, person.Age)
}
