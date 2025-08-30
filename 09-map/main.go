package main

import "fmt"

func main() {
	fmt.Println("Belajar Map")

	// Deklarasi map
	umur := map[string]int{
		"Kevin": 19,
		"Rizal": 20,
		"Adit":  21,
	}

	// Akses value
	fmt.Println("Umur Kevin:", umur["Kevin"])

	// Tambah data baru
	umur["Laura"] = 22
	fmt.Println("Setelah tambah Laura:", umur)

	// Update data
	umur["Adit"] = 25
	fmt.Println("Update umur Adit:", umur)

	// Hapus data
	delete(umur, "Rizal")
	fmt.Println("Setelah hapus Rizal:", umur)

	// Looping map
	for nama, usia := range umur {
		fmt.Println("Nama:", nama, "- Umur:", usia)
	}

	// Cek apakah key ada
	nilai, ada := umur["Rizal"]
	fmt.Println("Umur Rizal:", nilai, "Ada?", ada)
}
