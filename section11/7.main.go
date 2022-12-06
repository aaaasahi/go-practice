package main

import "fmt"

// struct 独自型

// 独自の型を定義
type MyInt int

// 独自の型を使ってメソッドを作る
func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi) // 0
	fmt.Printf("%T\n", mi) // main.MyInt

	mi.Print()
}
