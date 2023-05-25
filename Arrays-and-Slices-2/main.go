package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 12
	a[1] = 24
	a[2] = 32
	fmt.Println(a)
	// other print >>> [12 24 32]

	b := [3]int{12, 78, 90}
	fmt.Println(b)

	// other print >>> [12 78 90]

	c := [...]int{12, 78, 29}
	fmt.Println(c)

	//other print >>> [12 78 29]

	var d [3]int
	fmt.Println(d)

	//print >>> [0 0 0]

	e := [3]int{12}
	fmt.Println(e)

	//print >>> [12 0 0]
}
