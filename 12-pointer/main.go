package main

import "fmt"

// ===== 1) Dasar Pointer =====
func basicPointer() {
	fmt.Println("== Basic Pointer ==")
	x := 10
	px := &x // px menampung alamat memori x

	fmt.Println("x       :", x)
	fmt.Println("&x (px) :", px)
	fmt.Println("*px     :", *px) // dereference

	*px = 99 // ubah nilai lewat pointer
	fmt.Println("x setelah *px=99:", x)
	fmt.Println()
}

// ===== 2) Pointer sebagai parameter =====
func incrementByValue(v int) {
	v = v + 1 // hanya salinan, luar tidak berubah
}

func incrementByPointer(p *int) {
	*p = *p + 1 // mengubah nilai asli
}

func pointerParam() {
	fmt.Println("== Pointer Parameter ==")
	a := 5
	incrementByValue(a)
	fmt.Println("a setelah incrementByValue :", a) // tetap 5

	incrementByPointer(&a)
	fmt.Println("a setelah incrementByPointer:", a) // jadi 6
	fmt.Println()
}

// ===== 3) Pointer ke struct =====
type Student struct {
	Name   string
	Age    int
	School string
}

func pointerStruct() {
	fmt.Println("== Pointer Struct ==")
	s := Student{Name: "Kevin", Age: 19, School: "SMK N 1 Sumber"}

	ps := &s
	ps.Age = 20 // auto-deref, sama seperti (*ps).Age = 20

	fmt.Println("Struct s   :", s)
	fmt.Println("Pointer ps :", ps, "Age:", ps.Age)

	// pakai composite literal langsung pointer
	s2 := &Student{Name: "Adit", Age: 20, School: "SMK N 2 Pati"}
	fmt.Println("s2 pointer :", s2)
	fmt.Println()
}

// ===== 4) Method: value receiver vs pointer receiver =====
type Counter struct{ N int }

// tidak mengubah aslinya (by value / copy)
func (c Counter) IncByValue() { c.N++ }

// mengubah aslinya (by pointer)
func (c *Counter) IncByPointer() { c.N++ }

func receiverDemo() {
	fmt.Println("== Method Receiver ==")
	c := Counter{N: 0}
	c.IncByValue()
	fmt.Println("Setelah IncByValue  :", c.N) // 0, tidak berubah

	c.IncByPointer()
	fmt.Println("Setelah IncByPointer:", c.N) // 1, berubah
	fmt.Println()
}

// ===== 5) new, nil, dan catatan make =====
func newNilMake() {
	fmt.Println("== new / nil / make ==")

	// nil pointer
	var p *int
	if p == nil {
		fmt.Println("p == nil (belum menunjuk kemana-mana)")
	}

	// new(T) mengalokasikan zero value lalu mengembalikan *T
	num := new(int) // *int dengan nilai awal 0
	fmt.Println("*num (awal):", *num)
	*num = 7
	fmt.Println("*num (ubah):", *num)

	// catatan:
	// - new dipakai untuk tipe nilai (struct, int, dst) jika ingin pointer.
	// - make dipakai untuk slice, map, channel (bukan untuk struct/int).
	fmt.Println()
}

func main() {
	basicPointer()
	pointerParam()
	pointerStruct()
	receiverDemo()
	newNilMake()
}
