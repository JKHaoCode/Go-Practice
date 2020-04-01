package main

import (
	"fmt"
)

func modify(x int) {
	x = 300
}

func modifyPointer(x *int) {
	*x = 300
}

// pointer
func main() {
	fmt.Println("GO")

	var n = 100
	p := &n
	fmt.Println(&n)
	fmt.Printf("%T\n", p)
	fmt.Println(n)
	modify(n)
	fmt.Println(n)
	modifyPointer(&n)
	fmt.Println(n)
}
