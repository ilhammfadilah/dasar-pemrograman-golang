package main

import "fmt"

func main() {
	// inisialisasi nilai awal array
	var fruits = [4] string {"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlan element \t\t", len(fruits))
	fmt.Println("Value of element \t", fruits)

	// inisisalilasi nialai awal tanpa jumlah elemen
	var numbers = [...] int {2, 8, 6, 1}
	fmt.Println("Jumlan element \t\t", len(numbers))
	fmt.Println("Value of element \t", numbers)

	// array multidimensi
	var numbers1 = [2][3] int {[3] int {3,2,3}, [3] int {3,4,5}}
	var numbers2 = [2][3] int {{3,2,3}, {3,4,5}}
	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	// perulangan element array dengan keyword for
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("element %d : %s \n", i, fruits[i])
	}

	// perulangan element array dengan for - range
	for i, fruit := range fruits {
		fmt.Printf("element %d : %s \n", i, fruit)
	}

	// penggunaan variable _ dalam for -range
	for _, fruit := range fruits {
		fmt.Printf("element : %s \n", fruit)
	}

	// alokasi element array dengan keyword make
	var animals = make([]string, 2)
	animals[0] = "cat"
	animals[1] = "dog"

	fmt.Println(animals)
}