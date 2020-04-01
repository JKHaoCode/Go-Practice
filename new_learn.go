package main

import (
	"fmt"
)

func f(n int) (r int) {
	// r = 0
	defer func() {
		r += n
	}()

	var fo func()
	fo = func() {
		r += 2
	}
	defer fo()

	return
}

func main() {
	fmt.Println(f(3))
}
