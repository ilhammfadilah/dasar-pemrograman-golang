package main

import "fmt"

func main() {
	// deklarasi variabel dengan menggunakan var keyword
	var first_name string = "Ilham"
	var last_name string
	last_name = "Fadilah"

	/* 
		deklarasi tanpa menggunakan tipe data
		hanya bisa digunakan dalam blok fungsi
	*/
	middle_name := "Muhamad"

	// deklarasi multi variabel
	var tanggal_lahir, alamat string
	tanggal_lahir, alamat = "29", "Ciamis"
	var hobby, status string = "Bermain game", "Single"
	golongan_darah, agama := "A", "Islam"

	// deklarasi dengan menggunakan underscore variabel
	_ = "belajar golang"

	// deklarasi dengan menggunakan keyword new
	usia := new(string)

	fmt.Printf("Hello %s %s %s!\n", first_name, middle_name, last_name)
}