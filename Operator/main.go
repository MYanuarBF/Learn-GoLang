package main

import (
	"fmt"
)

func main() {
	fmt.Println(1 + 1)      //penjumlahan
	fmt.Println(10 - 1)     //pengurangan
	fmt.Println(10 * 2)     //perkalian
	fmt.Println(10 / 2)     //bagi
	fmt.Println(10.0 / 3.0) //bagi float
	fmt.Println(10 % 3)     //modulo

	// not expression
	fmt.Println(!false)
	fmt.Println(!true)
	//
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
}
