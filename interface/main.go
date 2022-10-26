package main

import (
	"fmt"
	"math"
	"strings"
)

type hitung interface {
	luas() float64
	keliling() float64
}

type cetak interface {
	getString() string
}

type mahasiswa struct {
	nama string
	jmlHuruf int
}

func (m mahasiswa) getString() string {
	return m.nama
}

func (m mahasiswa) getLength() int {
	return m.jmlHuruf
}

type lingkaran struct {
	diameter float64
}

func (l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func (l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func (l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}

type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}
func (p persegi) keliling() float64 {
	return p.sisi * 4
}

// embedded interface

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type menghitung interface {
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

// Casting Variabel Interface Kosong Ke Objek Pointer
type person struct {
	name string
	age	int
}

func main() {
	var bangunDatar hitung

	bangunDatar = persegi{10.0}
	fmt.Println("===== persegi")
	fmt.Println("luas:", bangunDatar.luas())
	fmt.Println("keliling:", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("===== lingkaran")
	fmt.Println("luas:", bangunDatar.luas())
	fmt.Println("keliling:", bangunDatar.keliling())
	fmt.Println("jari-jari :", bangunDatar.(lingkaran).jariJari()) // <-- belum paham

	// embedded interface
	var bangunRuang = &kubus{4}
	fmt.Println("===== kubus")
	fmt.Println("luas :", bangunRuang.luas())
	fmt.Println("keliling :", bangunRuang.keliling())
	fmt.Println("volume :", bangunRuang.volume())

	// interface kosong
	var secret interface{}
	secret = "ilham"

	fmt.Println(secret)

	secret = []string{"golang", "javascript", "php"}
	fmt.Println(secret)

	secret = [3]int{1, 2, 3}
	fmt.Println(secret)

	secret = lingkaran{}
	fmt.Println(secret)

	// pada go v1.18 interface kosong bisa menggunakan keyword any
	var data map[string]any
	data = map[string]any{
		"name":  "ilham",
		"grade": 9,
	}

	fmt.Println(data)

	secret = 2
	var nomor = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is", nomor)

	secret = []string{"apple", "orange", "banana"}
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")

	var tampilNama cetak
	tampilNama = mahasiswa{"ilham", 9}
	fmt.Println(tampilNama.getString())

	var rahasia interface{} = &person{name: "ilham", age: 22}
	var name = rahasia.(*person).name
	fmt.Println(name)

	// Kombinasi Slice, map , dan interface{}
	var people = []map[string]interface{}{
		{"name": "ilham", "age": 22},
		{"name": "muhamad", "age": 21},
		{"name": "fadilah", "age": 20},
	}

	for _, each := range people {
		fmt.Println(each["name"], " age is ", each["age"])
	}
}
