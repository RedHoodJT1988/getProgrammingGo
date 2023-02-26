package main

import "fmt"

func main() {
	var room = "cave"

	if room == "mountain" {
		fmt.Println("You continue your trek into the mountains.")
	} else if room == "entrance" {
		fmt.Println("You light your torch and venture inside the cave.")
	} else if room == "cave" {
		fmt.Println("You use your torch inside the cave to look around. You find a chest with a flare gun and two flares.")
	} else {
		fmt.Println("I'm sorry that is not an option. Choose again.")
	}
}
