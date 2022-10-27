package main

import (
	"fmt"
	"strconv"
)

func main() {
	// string -> int
	// var number = "29"

	// var num, err = strconv.Atoi(number)

	// if err == nil {
	// 	fmt.Println(num)
	// }

	// int -> string
	// var num = 124
	// var str = strconv.Itoa(num)
	// fmt.Println(str) // "124"

	/*
		Digunakan untuk konversi string berbentuk numerik dengan basis tertentu ke tipe numerik non-desimal dengan lebar data bisa
		ditentukan.
	*/
	// var str = "124"
	// var num, err = strconv.ParseInt(str, 10, 64)
	// if err == nil {
	// 	fmt.Println(num) // 124
	// }

	// int64 ke string

	// var num = int64(24)
	// var str = strconv.FormatInt(num, 8)
	// fmt.Println(str)

	// Digunakan untuk konversi string ke numerik desimal dengan lebar data bisa ditentukan.

	// var str = "24.12"
	// var num, err = strconv.ParseFloat(str, 32)
	// if err == nil {
	// 	fmt.Println(num) // 24.1200008392334
	// }

	/*
		Berguna untuk konversi data bertipe float64 ke string dengan format eksponen, lebar digit desimal, dan lebar tipe data bisa
		ditentukan.
	*/

	// var num = float64(24.12)
	// var str = strconv.FormatFloat(num, 'f', 6, 64)
	// fmt.Println(str)

	// Digunakan untuk konversi bool ke string.
	var bul = true
	var str = strconv.FormatBool(bul)
	fmt.Println(str)
}
