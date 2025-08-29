package main

import "fmt"

func main() {
	// Constant (tidak bisa diubah nilainya)
	const pi = 3.14
	const appName = "Learn Go"

	// Multiple constants
	const (
		country = "Indonesia"
		year    = 2025
	)

	fmt.Println("Pi:", pi)
	fmt.Println("App:", appName)
	fmt.Println("Country:", country)
	fmt.Println("Year:", year)
}
