package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(345000001) + 56000000
	fmt.Println(num)
}