package main

import (
	"fmt"
	"strings"
)

func main() {
	var isExists = strings.Contains("john wick", "wick")
	fmt.Println(isExists)

	// Digunakan untuk deteksi apakah sebuah string (parameter pertama) diawali string tertentu (parameter kedua).
	var isPrefix1 = strings.HasPrefix("john wick", "jo")
	fmt.Println(isPrefix1)
	var isPrefix2 = strings.HasPrefix("john wick", "wi")
	fmt.Println(isPrefix2)

	// Digunakan untuk deteksi apakah sebuah string (parameter pertama) diakhiri string tertentu (parameter kedua).
	var isSuffix1 = strings.HasSuffix("john wick", "ic")
	fmt.Println(isSuffix1)
	var isSuffix2 = strings.HasSuffix("john wick", "ck")
	fmt.Println(isSuffix2)

	// Memiliki kegunaan untuk menghitung jumlah karakter tertentu (parameter kedua) dari sebuah string (parameter pertama).
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany)

	// Digunakan untuk mencari posisi indeks sebuah string (parameter kedua) dalam string (parameter pertama).
	var index1 = strings.Index("ethan hunt", "ha")
	fmt.Println(index1)
	/*
	   Fungsi ini digunakan untuk replace atau mengganti bagian dari string dengan string tertentu. Jumlah substring yang di-replace bisa
	   ditentukan, apakah hanya 1 string pertama, 2 string, atau seluruhnya.
	*/
	var text = "banana"
	var find = "a"
	var replaceWith = "o"
	var newText1 = strings.Replace(text, find, replaceWith, 1)
	fmt.Println(newText1) // "bonana"
	var newText2 = strings.Replace(text, find, replaceWith, 2)
	fmt.Println(newText2) // "bonona"
	var newText3 = strings.Replace(text, find, replaceWith, -1)
	fmt.Println(newText3) // "bonono"

	// Digunakan untuk mengulang string (parameter pertama) sebanyak data yang ditentukan (parameter kedua).
	var str = strings.Repeat("na", 4)
	fmt.Println(str)

	/*
			Digunakan untuk memisah string (parameter pertama) dengan tanda pemisah bisa ditentukan sendiri (parameter kedua). Hasilnya
		berupa slice string.
	*/
	var string1 = strings.Split("the dark knight", " ")
	fmt.Println(string1) // output: ["the", "dark", "knight"]
	var string2 = strings.Split("batman", "")
	fmt.Println(string2) // output: ["b", "a", "t", "m", "a", "n"]

	/*
			Digunakan untuk menggabungkan slice string (parameter pertama)
		menjadi sebuah string dengan pemisah tertentu (parameter kedua.
	*/
	var data = []string{"banana", "papaya", "tomato"}
	var str1 = strings.Join(data, "-")
	fmt.Println(str1) // "banana-papaya-tomato"

	var str3 = strings.ToLower("aLAy")
	fmt.Println(str3)

	var str4 = strings.ToUpper("kecil")
	fmt.Println(str4)
}
