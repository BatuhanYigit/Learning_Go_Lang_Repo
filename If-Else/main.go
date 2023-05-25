package main

import "fmt"

func main() {
	num := 1
	if num%2 == 0 {
		fmt.Println("The Number ", num, "is even")
	} else {
		fmt.Println("The Number", num, "is odd")
	}
}
