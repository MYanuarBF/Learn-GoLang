package main

import "fmt"

func main() {
	var nama = [4]string{"Muhammad", "yanuar"}
	// menambah array
	nama[2] = "Budi"
	nama[3] = "Fatmadi"

	//penulisan cara lama/manual
	for i := 0; i < len(nama); i++ {
		fmt.Println(nama[i])
	}

	//penulisan cara baru/lebih efektif
	for index, value := range nama {
		fmt.Println("index", index, "=", value)
		// fmt.Println(value) //jika ingin menampilkan value saja ubah index menjadi '_'.
	}
}
