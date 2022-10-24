package main

import {
	"fmt"
	"strings"
	"math/rand"
	"time"
}

func main() {
	// penerapan fungsi
	var names = []string{"John", "Wick"}
	// printMessage("halo", names)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

// fungsi dengan return value