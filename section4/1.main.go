package main

import "fmt"

// int型
func main() {
	var i int = 100

	var i2 int64 = 200

	fmt.Println(i + 50)

	// 型が違うので実行できない
	// fmt.Println(i + i2)

	// 型の表示
	fmt.Printf("%T\n", i)

	// 型変換
	fmt.Printf("%T\n", int32(i2))

	// 型を変換することで実行できる
	fmt.Println(i + int(i2))
}