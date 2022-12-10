package main

// スコープ

import (
	"fmt"
	// f "fmt" fで使用できる（f.Println）
	"go-practice/section13/foo"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	b = s
	{
		b := "BBBB"
		fmt.Println(b) // BBBB この中でのスコープ
	}
	fmt.Println(b) // AAA
	return b
}

func main() {
	fmt.Println(foo.Max) // 100
	// fmt.Println(foo.min) // 呼び出せない

	fmt.Println(foo.ReturnMin()) // 1 関数の返り値の場合はprivateでも呼び出せる

	fmt.Println(appName()) // GoApp 1.0
	// fmt.Println(AppName, Version) // appName内の定数、変数は呼び出せない

	fmt.Println(Do("AAA"))
}
