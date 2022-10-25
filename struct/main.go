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

}

type student struct {
	name string
	grade int
	person
}

type person struct {
	age int
}
