package main

import (
	"fmt"
)

func product(out chan<- int) {
	// defer close(out)

	for i := 0; i < 5; i++ {
		out <- i
	}
}

func cunsumer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)
	go product(ch)
	cunsumer(ch)

	fmt.Println("main 结束")
}
