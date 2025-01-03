package main

import "fmt"

func main() {

	var nama string = "Yanuar Budi"
	fmt.Println("Hello, Yanuar")
	fmt.Println("Hello, " + nama + " Apa Kabar?")
	fmt.Println(len("acumalaka")) // menghitung panjang string
	fmt.Println("GOLANG"[0])      //mengambil char dari string
	fmt.Println("GOLANG"[0:4])    //mengambil beberapa char dari string
	fmt.Println("GOLANG"[3:])     //mengambil beberapa char dari string mulai dari akhir
	fmt.Println("GOLANG"[:3])     //mengambil beberapa char dari string mulai dari awal
}
