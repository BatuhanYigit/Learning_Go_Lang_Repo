package main

import (
	"fmt"
)

func calculater(price, no int) int {
	var totalprice = price * no
	return totalprice

}

func main() {
	var price, no = 10, 20
	totalprice := calculater(price, no)
	fmt.Println("Total Price : ", totalprice)
}
