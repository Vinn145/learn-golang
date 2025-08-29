package main

import "fmt"

func main() {
	// Deklarasi variabel dengan tipe data eksplisit
	var name string = "Kevin"
	var age int = 19

	// Deklarasi variabel tanpa tipe (tipe otomatis)
	city := "Pati"
	score := 95.5

	// Multi-deklarasi
	var x, y int = 10, 20

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("City:", city)
	fmt.Println("Score:", score)
	fmt.Println("x + y =", x+y)
}
