package main

import "fmt"

func main() {
	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and 50", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less 100", num)
	}
}
