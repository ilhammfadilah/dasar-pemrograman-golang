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
		fmt.Printf("maaf tidak lulus, nilai anda %d \n", point)
	}

	// selekse if - else variabel temporary
	var skor = 8840.0
	if percent := skor / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad! \n", percent, "%")
	}

}
