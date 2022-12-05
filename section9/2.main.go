package main

import "fmt"

// スライス
// append make len cap
func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	// 要素の追加
	// 配列と違いスライスは要素の追加ができる
	sl = append(sl, 300)
	fmt.Println(sl) // [100 200 300]

	// 指定した要素数を作成
	sl2 := make([]int, 5)
	fmt.Println(sl2) // [0 0 0 0 0]

	// 要素数を出力
	fmt.Println(len(sl2)) // 5

	// 容量
	fmt.Println(cap(sl2)) // 5

	sl3 := make([]int, 5, 10) // 第3引数で容量を指定
	fmt.Println(len(sl3)) // 5
	fmt.Println(cap(sl3)) // 10
}