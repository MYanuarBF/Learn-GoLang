package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		if i == 50 {
			break // stop program
		}
	}

	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			continue                 // skip kode dibawahnya
			fmt.Println("acumalaka") // tidak akan tampil karena di skip
		}
	}
}
