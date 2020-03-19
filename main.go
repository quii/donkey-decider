package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	donkeys := []string{
		"Osh",
		"Lisa",
		"Riya",
		"Tamara",
		"Vanessa",
		"Dave",
		"Chris",
		"Ivo",
		"Mikail",
		"Martyna",
		"Drilon",
	}

	rand.Seed(time.Now().Unix())

	selectedDonkey := donkeys[rand.Intn(len(donkeys))]

	fmt.Println("Today", selectedDonkey, "will be running the standup")
}