package main

import "fmt"

func main() {
	// data pribadi
	name := "Kevin Andrea"
	age := 19
	graduated := true
	school := "SMK 1 Sumber"
	district := "Batangan"
	village := "Kuniran"

	fmt.Println("=== Profil Singkat ===")
	fmt.Println("Nama:", name)
	fmt.Println("Umur:", age)
	fmt.Println("Sekolah:", school)
	fmt.Println("Sudah lulus:", graduated)

	// contoh IF
	if age < 13 {
		fmt.Println("Status umur: Anak-anak")
	} else if age < 20 {
		fmt.Println("Status umur: Remaja")
	} else {
		fmt.Println("Status umur: Dewasa")
	}

	// contoh SWITCH untuk lokasi
	switch district {
	case "Pati":
		fmt.Println("Lokasi: Kabupaten Pati")
	case "Batangan":
		fmt.Println("Lokasi: Kecamatan Batangan")
	default:
		fmt.Println("Lokasi: Kecamatan lain")
	}

	switch village {
	case "Kuniran":
		fmt.Println("Tinggal di Desa Kuniran")
	default:
		fmt.Println("Desa tidak dikenal")
	}

	// contoh gabungan IF + SWITCH
	if graduated {
		fmt.Println("Status: Baru lulus,", school)
		switch age {
		case 18, 19, 20:
			fmt.Println("saya siap lanjut kuliah atau kerja 👍")
		default:
			fmt.Println("Umur tidak biasa untuk baru lulus SMK")
		}
	} else {
		fmt.Println("Masih sekolah di", school)
	}
}
