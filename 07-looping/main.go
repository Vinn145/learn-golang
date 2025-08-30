package main

import "fmt"

func main() {
	fmt.Println("Belajar Looping di Golang")

	// Loop standar (for i)
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
	}

	// Loop gaya while
	j := 1
	for j <= 3 {
		fmt.Println("Nilai j:", j)
		j++
	}

	// Loop range (untuk slice/array)
	angka := []int{10, 20, 30, 40}
	for index, value := range angka {
		fmt.Println("Index:", index, "Value:", value)
	}
}
