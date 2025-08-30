package main

import "fmt"

func main() {
	// Data personal
	var name string = "Kevin Andrea"
	var age int = 19
	var gpa float32 = 3.50      // contoh IPK
	var isGraduated bool = true // sudah lulus SMK

	city := "Pati"
	school := "SMK 1 Sumber"
	year := 2025

	// Output
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("GPA:", gpa)
	fmt.Println("Graduated:", isGraduated)
	fmt.Println("City:", city)
	fmt.Println("School:", school)
	fmt.Println("Year:", year)

	// Cek tipe data
	fmt.Printf("Type of gpa: %T\n", gpa)
	fmt.Printf("Type of city: %T\n", city)
	fmt.Printf("Type of graduated: %T\n", isGraduated)
}
