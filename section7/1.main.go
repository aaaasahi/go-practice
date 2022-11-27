package main

import "fmt"

// 関数

// 引数と返り値の型を指定
func Plus(x, y int) int {
	return x + y
}

// 複数の返り値の場合
func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

// 返り値の変数宣言
func Double(price int) (result int) {
	result = price * 2
	return
}

// 返り値のない関数
func Noreturn() {
	fmt.Println("No Return")
	return
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3)

	i4 := Double(1000)
	fmt.Println(i4)

	Noreturn()
}
