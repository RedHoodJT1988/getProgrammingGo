package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number = 47

	for {
		var n = rand.Intn(100) + 1
		if n < number {
			fmt.Printf("%v is toom small.\n", n)
		} else if n > number {
			fmt.Printf("%v is too big.\n", n)
		} else {
			fmt.Printf("You got it! %v\n", n)
			break
		}
	}
}
