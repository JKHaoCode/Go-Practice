package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Age       int    `json:"age"`
	Name      string `json:"name`
	Niubility bool   `json:"niubility"`
}

func main() {
	b := []byte(`{"age":18,"name": "Mr.Sun", "niubility":true}`)
	var p Person

	err := json.Unmarshal(b, &p)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(p)
}
