package main
import "level-akses/library"
import "fmt"

func main() {
	// library.SayHello("ilham")
	// library.introduce("ilham")

	var student1 = library.Student{"ilham", 9}
	fmt.Println("student 1 name: ", student1.Name)
	fmt.Println("student 1 grade: ", student1.Grade)
}