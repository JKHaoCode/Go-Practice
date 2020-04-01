package main

import (
	"encoding/json"
	"fmt"
)

type point struct {
	X int `json:"lisi"`
	Y int `json:"wangwu"`
}

func main() {
	str := `{"lisi": 10, "wangwu": 35}`

	var po point

	err := json.Unmarshal([]byte(str), &po)

	if err != nil {
		fmt.Println("unmarsha1 err is ", err)
	}

	fmt.Println(po)
}
