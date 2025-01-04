package main

import (
	"fmt"
	"strconv"
)

func main() {
	//String to Integer
	// using strconv
	nomorTelp := "087822308404"
	fmt.Println(nomorTelp)

	Telepon, _ := strconv.Atoi(nomorTelp)
	fmt.Println(Telepon)

	tilpun := strconv.Itoa(Telepon)
	fmt.Println(tilpun)

	//convert float to integer
	var IPK float32 = 3.82
	fmt.Println(IPK)

	var GPA int = int(IPK)
	fmt.Println(GPA)

}
