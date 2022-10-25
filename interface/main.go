package main

import "fmt"
import "math"

type hitung interface {
	luas() float64
	keliling() float64
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
}