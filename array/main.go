package main
import "fmt"

func main() {
	// inisialisasi nilai awal array
	var fruits = [4] string {"apple", "grape", "banana", "melon"}

	fmt.Println("Jumlan element \t\t", len(fruits))
	fmt.Println("Value of element \t", fruits)
}