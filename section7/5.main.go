package main

import "fmt"

// クロージャー
func Later() func(string) string {
	var store string // 初期値空文字
	return func(next string) string {
		s := store
		store = next // 引数をstoreに代入、次回実行時保持される
		return s
	}
}

func main() {
	f := Later
	fmt.Println(f("Hello")) // ' '
	fmt.Println(f("My"))    // Hello
	fmt.Println(f("name"))  // My
	fmt.Println(f("is"))    // name
	fmt.Println(f("Golang")) // is
}