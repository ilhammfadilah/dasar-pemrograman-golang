package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	defer catch()

	// var input string
	// fmt.Print("masukan nomor : ")
	// fmt.Scanln(&input)

	// var number int
	// var err error
	// number, err = strconv.Atoi(input)

	// if err == nil {
	// 	fmt.Println(number, " adalah nomor")
	// } else {
	// 	fmt.Println(input, " bukan nomor")
	// 	fmt.Println(err.Error())
	// }

	var nama string
	fmt.Print("masukan nama anda :")
	fmt.Scanln(&nama)
	if valid, err := validate(nama); valid {
		fmt.Println("halo ", nama)
	} else {
		// fmt.Println(err.Error())

		// panic
		panic(err.Error())
		fmt.Println("end")
	}
}

// membuat custom error
func validate(nama string) (bool, error) {
	if strings.TrimSpace(nama) == "" {
		return false, errors.New("tidak boleh kosong")
	}
	return true, nil
}

// recover
func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}
