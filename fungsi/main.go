package main

import (
	"fmt"
	"strings"
	"math"
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

	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran \t\t: %.2f \n", area)
	fmt.Printf("keliling lingkaran \t: %.2f \n", circumference)
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

// fungsi multiple return
func calculate(d float64) (float64, float64) {
	// hitung luas
	var area = math.Pi * math.Pow(d / 2, 2)
	// hitung keliling
	var circumference = math.Pi * d

	// kembalikan 2 nilai
	return area, circumference
}