package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var jsonString = `{"Name": "John Wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user: ", data.FullName)
	fmt.Println("age: ", data.Age)

	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("Name: ", data1["Name"])
	fmt.Println("Age: ", data1["Age"])

	var jsonString2 = `[
		{"Name": "ilham muhamad fadilah", "Age": 22},	
		{"Name": "agung nugraha", "Age": 32}	
	]`

	var data2 []User
	var error = json.Unmarshal([]byte(jsonString2), &data2)
	if error != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1: ", data2[0].FullName)
	fmt.Println("user 1: ", data2[0].Age)
}
