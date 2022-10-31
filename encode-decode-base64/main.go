package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// var data = "john wick"

	// var encodingString = base64.StdEncoding.EncodeToString([]byte(data))
	// fmt.Println("encode : ", encodingString)

	// var decodeByte, _ = base64.StdEncoding.DecodeString(encodingString)
	// var decodeString = string(decodeByte)
	// fmt.Println("decode : ", decodeString)

	// var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	// base64.StdEncoding.Encode(encoded, []byte(data))
	// var encodedString = string(encoded)
	// fmt.Println(encodedString)

	// var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	// var _, err = base64.StdEncoding.Decode(decoded, encoded)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// var decodedString = string(decoded)
	// fmt.Println(decodedString)

	var url = "https://kalipare.com/"
	var encodedString = base64.URLEncoding.EncodeToString([]byte(url))
	fmt.Println(encodedString)
	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)
}
