package main

import "fmt"

func main() {
	var deciman_number = 1.2 //tipe data desimal
	var exist = true // tipe data bool
	var message string = "Assalamualaikum" // tipe data string
	const pi = 3.14 // tipe data konstanta

	fmt.Printf("Bilangan desimal: %f\n", deciman_number);
	fmt.Printf("Bilangan desimal: %.2f\n", deciman_number);
	
	fmt.Printf("Exist: %t \n", exist);

	fmt.Printf("Salam: %s \n", message);
}