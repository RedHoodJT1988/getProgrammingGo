// How long does it take to get to Mars?
package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	var distance = 56000000   // km

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	const hoursPerDay, minutesPerHour = 24, 60
	var speed = 100800      // km/h
	var distance2 = 96300000 // km

	var weight = 220
	weight -= 2

	fmt.Println(distance2/speed/hoursPerDay, "days")
}