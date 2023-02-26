package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")

	var command = "walk outside"
	var exit = strings.Contains(command, "outside")

	fmt.Println("You leave the cave:", exit)

	var wearShades = true
	fmt.Println("You are wearing shades:", wearShades)

	fmt.Println("There is a sign near the cave entrance.")
	var read = strings.Contains(command, "read")
	fmt.Println(read)
}
