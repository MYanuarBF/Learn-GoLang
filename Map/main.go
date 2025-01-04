package main

import "fmt"

func main() {
	mahasiswa := make(map[string]string)
	mahasiswa["2272"] = "Yanuar"
	mahasiswa["2200"] = "Vangga"
	mahasiswa["2201"] = "Varrel"

	fmt.Println(mahasiswa)

	for nim, nama := range mahasiswa {
		fmt.Println("NIM", nim, "Bernama", nama)
	}

	//map didalam map
	datadiri := map[string]map[string]string{
		"331": map[string]string{
			"nama":   "Yanuar",
			"alamat": "Bumiayu",
		},
		"332": map[string]string{
			"nama":   "Vangga",
			"alamat": "Banjarnegara",
		},
		"333": map[string]string{
			"nama":   "Varrel",
			"alamat": "Purbalingga",
		},
	}

	// menghapus data pada MAP
	delete(datadiri, "333")

	//menampilkan MAP
	for _, dadir := range datadiri {
		fmt.Println(dadir)
	}
}
