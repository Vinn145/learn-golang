package main

import "fmt"

func main() {
	fmt.Println("=== Arithmetic Operators ===")
	arithmetic()

	fmt.Println("\n=== Relational Operators ===")
	relational()

	fmt.Println("\n=== Logical Operators ===")
	logical()

	fmt.Println("\n=== Assignment Operators ===")
	assignment()
}

func arithmetic() {
	a, b := 10, 3
	fmt.Println("a + b =", a+b)
	fmt.Println("a - b =", a-b)
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)
	fmt.Println("a % b =", a%b)
}

func relational() {
	a, b := 5, 10
	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a <= b:", a <= b)
	fmt.Println("a >= b:", a >= b)
}

func logical() {
	x, y := true, false
	fmt.Println("x && y:", x && y)
	fmt.Println("x || y:", x || y)
	fmt.Println("!x:", !x)
}

func assignment() {
	num := 10
	fmt.Println("num =", num)
	num += 5
	fmt.Println("num += 5 ->", num)
	num -= 3
	fmt.Println("num -= 3 ->", num)
	num *= 2
	fmt.Println("num *= 2 ->", num)
	num /= 4
	fmt.Println("num /= 4 ->", num)
	num %= 3
	fmt.Println("num %= 3 ->", num)
}
