package main

import "fmt"

// byte型
func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	// byteをstringに変換
	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)

	// stringをbyteに変換
	fmt.Println(string(c))
}