package main

import "fmt"

func main() {
	a := [...]string{"ABD", "TURKEY", "CHINA", "INDIA", "GERMANY"}
	b := a
	b[0] = "Singapore"
	fmt.Println("a is", a)
	fmt.Println("b is", b)

	//change data

}
