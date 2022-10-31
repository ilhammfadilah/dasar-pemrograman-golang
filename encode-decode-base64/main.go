package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var data = "john wick"

	var encodingString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encode : ", encodingString)

	var decodeByte, _ = base64.StdEncoding.DecodeString(encodingString)
	var decodeString = string(decodeByte)
	fmt.Println("decode : ", decodeString)
}
