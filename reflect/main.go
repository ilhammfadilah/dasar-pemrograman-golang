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

	// Pengaksesan Nilai Dalam Bentuk interface{}
	fmt.Println("nilai variabel : ", reflectValue.Interface())

	var nilai = reflectValue.Interface().(int)
	fmt.Println(nilai)

	// Pengaksesan Informasi Property Variabel Objek
	var student1 = &student{Name: "Ilham", Grade: 9}
	student1.getPropertyInfo()

	// Pengaksesan Informasi Method Variabel Objek
	var student2 = &student{Name: "fadilah", Grade: 8}
	fmt.Println("student 2 name: ", student2.Name)

	var reflectValueStudent2 = reflect.ValueOf(student2)
	var method = reflectValueStudent2.MethodByName("SetName") // mengambil method SetName
	method.Call([]reflect.Value{ // method dipanggil, parameter harus berbentuk array berurutan sesuai urutan dari deklari parameternya
		reflect.ValueOf("moona"), // merubah value dari parameter pertama
	})
	fmt.Println("nama : ", student2.Name)
}

type student struct {
	Name string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s) // << object reflect.Value dari variabel s
	if reflectValue.Kind() == reflect.Ptr { // << check apakah variabel tersebut pointer
		reflectValue = reflectValue.Elem() // << mengambil object reflect aslinya
	}

	var reflectType = reflectValue.Type() // informasi type
	for i := 0; i < reflectValue.NumField(); i++ { // << perulangan sebanyak jumlah properti dalam struct
		fmt.Println("nama		:", reflectType.Field(i).Name)
		fmt.Println("tipe data	:", reflectType.Field(i).Type)
		fmt.Println("nilai		:", reflectValue.Field(i).Interface())
		fmt.Println("")
	}
}

func (s *student) SetName(name string) {
	s.Name = name
}