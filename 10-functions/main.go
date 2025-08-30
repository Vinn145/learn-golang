package main

import (
	"fmt"
)

// 1. Function tanpa parameter
func sayHello() {
	fmt.Println("Hello, selamat datang di Golang!")
}

// 2. Function dengan parameter
func greet(name string) {
	fmt.Println("Halo,", name)
}

// 3. Function dengan return value
func add(a int, b int) int {
	return a + b
}

// 4. Function dengan multiple return
func divide(dividend, divisor int) (int, int) {
	hasil := dividend / divisor
	sisa := dividend % divisor
	return hasil, sisa
}

func main() {
	// 1. Panggil function tanpa parameter
	sayHello()

	// 2. Panggil function dengan parameter
	greet("Kevin Andrea")

	// 3. Panggil function dengan return value
	sum := add(5, 3)
	fmt.Println("Hasil penjumlahan:", sum)

	// 4. Panggil function dengan multiple return
	hasil, sisa := divide(10, 3)
	fmt.Println("Hasil pembagian:", hasil, "Sisa:", sisa)

	// 5. Anonymous function (langsung dipanggil)
	func(msg string) {
		fmt.Println("Pesan dari anonymous function:", msg)
	}("Halo dari anonymous!")

	// 6. Variadic function (bisa input banyak angka)
	total := sumAll(1, 2, 3, 4, 5)
	fmt.Println("Total variadic sum:", total)
}

// 6. Variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
