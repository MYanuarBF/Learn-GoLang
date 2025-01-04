package main

import "fmt"

func main() {
	// melakukan looping
	fmt.Println("Looping 'hello Yanuar' 10x")
	i := 1

	for i <= 10 {
		fmt.Println(i, "Hello Yanuar")
		i++
	}
	fmt.Println("Selesai")

	// cara penulisan alternatif
	for i := 10; i >= 1; i-- {
		fmt.Println(i, "Reverse loop")
	}
}
