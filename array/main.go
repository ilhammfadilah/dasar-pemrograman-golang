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

}