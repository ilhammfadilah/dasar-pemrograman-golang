package main
import "fmt"

func main() {
	
	// perulangan dengan menggunakna keyword for
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// perulangan keyword for dengan argumen hanya kondisi
	var n = 0
	for n < 5 {
		fmt.Println("Angka", n)
		n++
	}
}