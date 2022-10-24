// Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair.

package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("februari", chicken["februari"])

	// inisialisasi nilai map
	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	fmt.Println(chicken1["februari"])

	// Iterasi Item Map Menggunakan for - range
	var egg = map[string]int{
		"januari": 10,
		"februari": 20,
		"maret": 30,
		"april": 40,
	}

	for key, val := range egg {
		fmt.Println(key, "\t :", val)
	}

	// Menghapus Item Map
	delete(egg, "april")
	fmt.Println(egg)
}