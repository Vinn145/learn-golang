package main

import "fmt"

func main() {
	fmt.Println("Belajar Array & Slice")

	// Array (panjang tetap)
	var angka [3]int
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30
	fmt.Println("Array angka:", angka)

	// Slice (lebih fleksibel)
	buah := []string{"Apel", "Mangga", "Jeruk"}
	fmt.Println("Slice buah:", buah)

	// Menambah elemen ke slice
	buah = append(buah, "Pisang")
	fmt.Println("Slice setelah ditambah:", buah)

	// Loop slice dengan range
	for i, v := range buah {
		fmt.Println("Index:", i, "Value:", v)
	}
}
