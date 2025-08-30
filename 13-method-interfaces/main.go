package main

import "fmt"

// Struct
type Person struct {
	Name string
	Age  int
}

// Method tanpa pointer receiver
func (p Person) SayHello() {
	fmt.Println("Halo, nama saya", p.Name, "dan umur saya", p.Age, "tahun.")
}

// Method dengan pointer receiver (bisa ubah nilai asli)
func (p *Person) Birthday() {
	p.Age++
	fmt.Println("Selamat ulang tahun,", p.Name, "sekarang umurmu", p.Age)
}

// Interface
type Info interface {
	GetInfo() string
}

// Implementasi interface pada Person
func (p Person) GetInfo() string {
	return fmt.Sprintf("Nama: %s, Umur: %d", p.Name, p.Age)
}

func main() {
	// Buat object Person
	kevin := Person{Name: "Kevin Andrea", Age: 18}

	// Panggil method
	kevin.SayHello()
	kevin.Birthday()

	// Pakai interface
	var i Info = kevin
	fmt.Println("Info lewat interface:", i.GetInfo())
}
