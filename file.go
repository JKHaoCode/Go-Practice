package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// file, err := os.Create("./xxddd.txt")

	// if err != nil {
	// 	fmt.Println("err", err)
	// 	return
	// }

	// defer file.Close()

	// for i := 0; i < 5; i++ {
	// 	file.WriteString("dfasdfasd\n")
	// 	file.Write([]byte("sdfasdfwwoo\n"))
	// }
	file, err := os.Open("./xxddd.txt")
	if err != nil {
		fmt.Println("open file err", err)
		return
	}

	defer file.Close()

	var buf [128]byte

	for {
		n, err := file.Read(buf[:])

		if err == io.EOF {

		}
	}
}
