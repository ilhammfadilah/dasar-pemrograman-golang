package main

import "fmt"

func main() {

	// seleksi kondisi menggunakan if, else if, dan else
	var point = 8

	if(point == 10) {
		fmt.Println("lulus dengan nilai sempurana")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("maaf tidak lulus, nilai anda %d", point)
	}
}
