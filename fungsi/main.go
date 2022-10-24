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

	divideNumber(10, 2)
	divideNumber(4, 0)
	divideNumber(8, -4)
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

// Penggunaan Keyword return Untuk Menghentikan Proses Dalam Fungsi
func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d \n", m, n)
		return
	}
	var res = m / n
	fmt.Printf("%d / %d = %d \n", m, n, res)
}