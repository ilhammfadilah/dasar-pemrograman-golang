package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int32
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.5,
	age:         29,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

func main() {
	// format biner
	fmt.Printf("%b \n", data.age)

	// format unicode
	fmt.Printf("%c \n", 1400)

	// format basis 10
	fmt.Printf("%d \n", data.age)

	// format notasi numerik standar
	fmt.Printf("%e \n", data.height)

	// format desimal dengan dengan lebar desimal bisa ditentukan
	fmt.Printf("%f \n", data.height)

	// format desimal dengan lebar
	fmt.Printf("%g \n", 0.123123123)

	// numberic -> octal
	fmt.Printf("%o \n", data.age)

	// format pointer
	fmt.Printf("%p \n", &data.name)

	// escape string
	fmt.Printf("%q \n", `"name \ height"`)

	// data string
	fmt.Printf("%s \n", data.name)

	// data boolean
	fmt.Printf("%t \n", data.isGraduated)

	// megambil tipe variabel
	fmt.Printf("%T \n", data.age)

	// format apa saja
	fmt.Printf("%v \n", data)

	// format struct
	fmt.Printf("%+v \n", data)

	// format sesuai struct
	fmt.Printf("%#v \n", data)

	// numerik -> string basis16
	fmt.Printf("%x \n", data.age)

	// menulis %
	fmt.Printf("%% \n")
}
