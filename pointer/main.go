/*
	pointer adalah reference atau alamat memori.
	variabel pointer berisi alamat memori suatu nilai.
	contoh sebuah variabel bertipe integer memiliki nilai 9, maka yang dimaksud ponter adalah alamat memori dimana nilai 9 itu disimpan.
	variabel - variabel yang memiliki reference atau alamat pointer yang sama, maka saling terhubung satu sama lain.
*/
package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("variable A (value): ", numberA)
	fmt.Println("variable A (address): ", &numberA)

	fmt.Println("vairable B (value):", *numberB)
	fmt.Println("vairable B (address):", numberB)
}