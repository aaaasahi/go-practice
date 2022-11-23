package main

import "fmt"

// 関数外での定義
// i5 := 500

// こっちは実行できる
// var i5 int = 500

func main() {
	// 明示的な定義
	// var 変数名 型 = 値

	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	// 初期値が入る
	var i3 int
	var s3 string
	fmt.Println(i3, s3)

	// 値更新
	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	// 暗黙的な定義
	// 変数名 := 値
	// 型推論、関数外での定義はできない
	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 再定義できない
	// i4 := 500
	// fmt.Println(i4)

	// 関数外での定義
	// fmt.Println(i5)

	// 使用していない変数がある場合はエラーを吐く
	// var s5 string = "Not Use"
}