package main

import (
	"fmt"
)

func main() {
	ss := "123454321"

	for i, v := range ss {
		fmt.Println(v)
		if ss[i] != ss[len(ss)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}

	fmt.Println("是回文")
}
