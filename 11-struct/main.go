package main

import "fmt"

// Definisi struct
type Student struct {
	Name   string
	Age    int
	School string
}

// Method (fungsi yang nempel ke struct Student)
func (s Student) SayHello() {
	fmt.Printf("Halo, nama saya %s, umur %d, sekolah di %s\n", s.Name, s.Age, s.School)
}

func main() {
	// Membuat struct
	s1 := Student{
		Name:   "Kevin",
		Age:    19,
		School: "SMK N 1 Sumber",
	}

	s2 := Student{
		Name:   "Adit",
		Age:    20,
		School: "SMA N 3 Rembang",
	}

	// Cetak data biasa
	fmt.Println("Data Student 1:", s1)
	fmt.Println("Data Student 2:", s2)

	// Panggil method
	s1.SayHello()
	s2.SayHello()
}
