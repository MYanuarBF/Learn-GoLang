package main

import "fmt"

func main() {
	nama := [5]string{
		"Muhammad",
		"Yanuar",
		"Budi",
		"Fatmadi",
		"Azman",
	}

	fmt.Println(nama)

	//data slice adalah referensi dari array asli
	//jika data diubah maka data pada array ikut berubah
	slice1 := nama[0:2]
	slice2 := nama[2:] // dari belakang
	slice3 := nama[:3] // dari depan
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	// membuat slice dengan array baru
	slice4 := make([]string, 5)
	copy(slice4, slice3)
	fmt.Println(slice4)

	// cara menambah data pada slices
	slice4[4] = "Vangga"
	slice4 = append(slice4, "Varrel")

	// nilai array pertama tidak berubah karena yang diubah adalah nilai array baru
	fmt.Println("===================")
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
}
