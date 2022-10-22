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

	// perulangan for tanpa argumen
	for {
		fmt.Println("Angka", n)

		n++
		if n == 29 {
			break
		}
	}

	// penggunaan break dan continue
	for i := 1; i <= 10; i++ {
		if i % 2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	// perulangan bersarang
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}