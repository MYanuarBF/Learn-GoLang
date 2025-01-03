package main

import "fmt"

func main() {
	// variable menggunakan var
	var greet string = "Hello!, "
	var name string = "Yanuar"
	fmt.Println(greet + name)

	name = "Budi"
	fmt.Println(greet + name)

	//atau bisa dengan menggunakan (:=). tipe data akan di define otomatis
	namaLengkap := "Muhammad Yanuar Budi Fatmadi"
	Usia := 23
	fmt.Println(namaLengkap + ", Usia anda sekarang adalah " + fmt.Sprint(Usia))

	//constanta menggunakan const(nilainya mutlak tidak bisa dirubah)
	const kampus string = "Telyu Purwokerto"
	fmt.Println(kampus)
}
