package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("selamat datang")
	fmt.Println("good bye")

	pilihMakanan("batagor")
	pilihMakanan("seblak")

	os.Exit(1)
}

func pilihMakanan(nama string) {
	defer fmt.Println("selamat datang di hem kedai")
	if nama == "seblak" {
		fmt.Print("selera yang bagus ", " ")
		fmt.Print("anda pasti suka pedas \n")
		return
	}
	fmt.Println("pilihan makanan anda :", nama)
}
