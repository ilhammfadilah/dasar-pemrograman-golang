/*
Slice adalah reference elemen array. Slice bisa dibuat, atau bisa juga dihasilkan dari manipulasi sebuah array ataupun slice lainnya.
Karena merupakan data reference, menjadikan perubahan data di tiap elemen slice akan berdampak pada slice lain yang memiliki
alamat memori yang sama.
*/

package main

import "fmt"

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println("fruits: ",fruits)
	fmt.Println("aFruits: ",aFruits)
	fmt.Println("bFruits: ",bFruits)

	fmt.Println("aaFruits: ",aaFruits)
	fmt.Println("baFruits: ",baFruits)

	baFruits[0] = "pinaple"

	fmt.Println("fruits: ",fruits)
	fmt.Println("aFruits: ",aFruits)
	fmt.Println("bFruits: ",bFruits)

	fmt.Println("aaFruits: ",aaFruits)
	fmt.Println("baFruits: ",baFruits)

	// fungsi len digunakan untuk menghitung jumlah element slice yang ada
	fmt.Println(len(fruits))

	// fungsi cap
	fmt.Println(cap(aFruits))
	fmt.Println(cap(bFruits))

	// fungsi append
	var cFruits = append(bFruits, "papaya")
	fmt.Println(cFruits)

	// fungsi copy
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "orange", "apple"}

	n := copy(dst, src)
	fmt.Println(n)

	dest := []string{
		"potato",
		"potato",
		"potato",
	}
	sour := []string{"watermelon", "pinnaple"}

	ncopy := copy(dest, sour)

	fmt.Println(dest)
	fmt.Println(sour)
	fmt.Println(ncopy)
}