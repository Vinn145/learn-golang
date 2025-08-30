package main

import (
	"fmt"

	"kevinandrea/hellopkg/greet"
)

func main() {
	greet.Hello("Kevin Andrea")

	message := greet.GetMessage("Pati")
	fmt.Println(message)
}
