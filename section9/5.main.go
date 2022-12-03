package main

import "fmt"

// スライス
// 可変長引数 引数の数を指定しないで渡せる

// スライスを引数として受け取れる
func Sum(s...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3)) // 6
	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7)) // 28
	fmt.Println(Sum()) // 0

  // スライスを引数として渡せる
	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))
}