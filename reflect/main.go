/*
	reflection adalah teknik untuk inspeksi variabel.
	mengambil nilai dari variabel tersebut atau memanipulasinya.
	cakupannya sangat luas.
	seperti melihat struktur variabel, tipe data, nilai pointer dan banyak lagi.
*/ 

package main

import (
	"fmt"
	"reflect"
	)

func main() {
	// Mencari Tipe Data & Value Menggunakan Reflect
	var number = 22
	var reflectValue = reflect.ValueOf(number)
	fmt.Println(reflectValue)
	fmt.Println("tipe variabel : ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel :", reflectValue.Int())
	}
}