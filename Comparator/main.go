package main

import "fmt"

func main() {
	a := 10
	b := 15
	// sama dengan
	hasil := a == b
	fmt.Println(hasil)
	// tidak sama dengan
	hasil = a != b
	fmt.Println(hasil)
	// kurang dari
	hasil = a < b
	fmt.Println(hasil)
	// lebih dari
	hasil = a > b
	fmt.Println(hasil)
	// kurang dari sama dengan

	a = 12
	b = 12

	hasil = a <= b
	fmt.Println(hasil)
	// lebih dari sama dengan
	hasil = a >= b
	fmt.Println(hasil)
}
