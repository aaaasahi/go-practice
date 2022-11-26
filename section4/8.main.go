package main

import (
	"fmt"
	"strconv"
)

// 型変換
func main() {
	// int → float
	var i int = 1
	fl64 := float64(i)
	fmt.Println(fl64)
	fmt.Printf("i = %T\n", i) // int
	fmt.Printf("fl64 = %T\n", fl64) // float64

	// float → int
	i2 := int(fl64)
	fmt.Printf("i2 = %T\n", i2) // int

	fl := 10.5
	i3 := int(fl)
	fmt.Printf("i3 = %T\n", i3) // int
	fmt.Println(i3) // 10 小数点以下は切り捨てになる

	// string → int
	var s string = "100"
	fmt.Printf("s = %T\n", s)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i) // => 100
	fmt.Printf("i = %T\n", i) // int

	// int → string
	var i4 int = 200
	s2 := strconv.Itoa(i4)
	fmt.Println(s2) // => 200
	fmt.Printf("s2 = %T\n", s2) // => string

	// string → byte
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b) // => [72 101 108 108 111 32 87 111 114 108 100]

	// byte → string
	h2 := string(b)
	fmt.Println(h2) // => Hello World
	fmt.Printf("h2 = %T\n", h2) // string
}