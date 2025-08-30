package main

import "fmt"

func main() {
	age := 19

	// if else if contoh umur
	if age < 13 {
		fmt.Println("Anak-anak")
	} else if age < 20 {
		fmt.Println("Remaja")
	} else {
		fmt.Println("Dewasa")
	}

	// switch dengan data lokasi nyata
	district := "Batangan"
	village := "Kuniran"

	switch district {
	case "Pati":
		fmt.Println("Kamu di Kabupaten Pati")
	case "Batangan":
		fmt.Println("Kamu di Kecamatan Batangan")
	default:
		fmt.Println("Kecamatan tidak dikenal")
	}

	switch village {
	case "Kuniran":
		fmt.Println("Kamu tinggal di Desa Kuniran")
	default:
		fmt.Println("Desa tidak dikenal")
	}
}
