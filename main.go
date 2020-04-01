package main

import (
	"fmt"
	"os"
	// "github.com/gin-gonic/gin"
	// "net/http"
)

func test01() {
	a := [5]int{1, 2, 3, 4, 5}
	a[1] = 123
	fmt.Println(a)
}

func getCireleArea(r float32) (are float32) {
	if r < 0 {
		panic("not fu")
	}
	are = r * r
	return
}

func test03() {

	defer func() {
		if err := recover(); err != nil { // recover() 获取panic
			fmt.Println(err)
		}
	}()
	getCireleArea(1)
}

func main() {
	test01()
	test03()
	var buf [16]byte
	os.Stdin.Read(buf[:])
	os.Stdout.WriteString(string(buf[:]))
	// fmt.Println(getCireleArea(-1))
	// fmt.Println("Hello World")

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "Hello, World\n")
	// })

	// r.GET("/user/:name", func(c *gin.Context) {
	// 	name := c.Param("name")
	// 	c.String(http.StatusOK, "Hello %s", name)
	// })

	// r.Run(":8080")
}
