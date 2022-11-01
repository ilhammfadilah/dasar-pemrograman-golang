package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	var text = "this is secret"
	// var sha = sha1.New()
	// sha.Write([]byte(text))

	// var encrypted = sha.Sum(nil)
	// var encryptedString = fmt.Sprintf("%x", encrypted)

	// fmt.Println(encryptedString)

	// Metode Salting Pada Hash SHA1
	fmt.Printf("original text : %s\n\n", text)

	var hashed1, salt1 = doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)

	_ = salt1

}

func doHashUsingSalt(text string) (string, string) {
	var salt = fmt.Sprintf("%d", time.Now().UnixNano())
	var saltedText = fmt.Sprintf("text : '%s', salt : %s", text, salt)
	fmt.Println(saltedText)
	var sha = sha1.New()
	sha.Write([]byte(saltedText))
	var encrypted = sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}
