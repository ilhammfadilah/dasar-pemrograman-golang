/*
	struct merupakan kumpulan property dan function yang dibungkus kedalam tipe data baru dengan nama tertentu.
*/

package main

import "fmt"

func main() {
	var s1 student
	s1.name = "ilham"
	s1.grade = 9

	var p2 = person{age: 9}
	var s2 = student{"muhamad", 10, p2}

	fmt.Println("name :", s1.name)
	fmt.Println("grade :", s1.grade)

	fmt.Println("name :", s2.name)
	fmt.Println("grade :", s2.grade)

	// variabel objek pointer
	var s3 = student{name: "fadilah", grade: 9}
	var s4 *student = &s3
	fmt.Println("student 3 name: ", s3.name)
	fmt.Println("student 4 name: ", s4.name)

	s3.name = "agung"
	fmt.Println("student 3 name: ", s3.name)
	fmt.Println("student 4 name: ", s4.name)

	// embedded struct
	var s5 = student{}
	s5.name = "mona <3"
	s5.age = 22
	s5.grade = 9

	fmt.Println("student 5 name: ",s5.name)
	fmt.Println("student 5 age: ",s5.person.age)
	fmt.Println("student 5 grade: ",s5.grade)

	var handphone1 = struct {
		handphone
		discount int
	}{} // <-- anonymous struct

	handphone1.handphone = handphone{brand: "iPhone", price: 500}
	handphone1.discount = 99

	fmt.Println("brand handphone :", handphone1.brand)
	fmt.Println("price handphone :", handphone1.price)
	fmt.Println("discount handphone :", handphone1.discount)

	// kombinasi slice dengan struct
	var allHandphone = []handphone{
		{brand: "Google Pixel", price: 100},
		{brand: "iPhone", price: 200},
		{brand: "Samsung", price: 300},
	}

	for _, hp := range allHandphone {
		fmt.Println(hp.brand, " harganya ", hp.price)
	}

	// inisialisasi slice anonymous struct
	var pegawai = []struct{
		person
		gaji int
	}{
		{person: person{22}, gaji: 5},
		{person: person{21}, gaji: 4},
		{person: person{20}, gaji: 3},
	}

	for _, pegawai := range pegawai {
		fmt.Println("usia pegawai", pegawai.age, " gaji: ", pegawai.gaji)
	}
}

type student struct {
	name string
	grade int
	person // <-- embedded struct
}

type person struct {
	age int
}

type handphone struct {
	brand string
	price int
}
