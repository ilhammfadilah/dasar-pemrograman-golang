package main

import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)

func main() {
	// penerapan fungsi
	// var names = []string{"John", "Wick"}
	// printMessage("halo", names)

	rand.Seed(time.Now().Unix())
	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number :", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number :", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number :", randomValue)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// fungsi dengan return value
func randomWithRange(min, max int) int {
	var value = rand.Int() % (max - min + 1) + min
	return value
}