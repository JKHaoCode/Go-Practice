package main

import (
	"fmt"
)

func f1() {
	fmt.Println("Hello shahe")
}

func f2(x int, y ...int) int {
	sun := x
	for _, v := range y {
		sun += v
	}
	return sun
}

func main() {
	fmt.Println("Go")
	f1()
	f1()
	f1()
	fmt.Println(f2(2, 3, 4, 5, 444, 22, 44))
}
