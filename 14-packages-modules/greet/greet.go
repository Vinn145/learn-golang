package greet

import "fmt"

// Fungsi untuk menyapa seseorang
func Hello(name string) {
	fmt.Println("Halo", name, "selamat datang di Go!")
}

// Fungsi yang mengembalikan string
func GetMessage(city string) string {
	return "Salam hangat dari " + city
}
