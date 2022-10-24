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

	var avg = calculateVariadic(2,4,3,5,4,3,3,5,5,3)
	var msg = fmt.Sprintf("rata - rata : %.2f", avg)
	fmt.Println(msg)

	var numbers = []int{2,4,3,5,4,3,3,5,5,3}
	var ratarata = calculateVariadic(numbers...)
	var pesan = fmt.Sprintf("rata - rata : %.2f", ratarata)
	fmt.Println(pesan)

	yourHobbies("ilham", "sleeping", "eating")

/*
	CLOSURE
	closure adalah sebuah fungsi yang bisa disimpan kedalam variabel. 
	closeure merupakan anonymous function. bisa dimanfaatkan untuk membungkus
	suatu proses yang hanya digunakan satu kali atau dalam blok tertentu saja.
*/

	var getMinMax = func(n []int)(int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\n min : %v\n max : %v\n", numbers, min, max)

/*
	Imediately-Invoke Function Expression (IIFE)
	closure ini di eksekusi langsung saat deklarasinya. bisa digunakan untuk membungkus proses yang hanya dilakukan satu sesekali, bisa mengembalikan nilai, bisa juga tidak.
*/

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(3)

	fmt.Println("original numbers :", numbers)
	fmt.Println("filtered numbers :", newNumbers)

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

// fungsi variadic atau pembuatan fungsi dengan parameter sejenis yg tak terbatas
func calculateVariadic(numbers ...int) float64 {
	var total int = 0
	for _, number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, my name is: %s \n", name)
	fmt.Printf("My hobbies are: %s \n", hobbiesAsString)
}




