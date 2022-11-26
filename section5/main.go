package main

import "fmt"

// 定数

// 頭文字大文字で他パッケージから呼び出せる
const Pi = 3.14

// 複数定義
const (
	URL = "http://xxx.co.jp"
	SiteName = "test"
)

// 値の省略
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

// 連続する値の定義
const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi) // 3.14

	fmt.Println(URL, SiteName) // http://xxx.co.jp test

	fmt.Println(A, B, C, D, E, F) // 1 1 1 A A A

	fmt.Println(c0, c1, c2) // 0 1 2
}