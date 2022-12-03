package main

import "fmt"

// copy
func main() {
	sl := []int{100, 200}
	sl2 := sl
	sl2[0] = 1000
	// 元のslも置き換わってしまう
	fmt.Println(sl, sl2) // [1000 200] [1000 200]

	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := make([]int, 5, 10)
	fmt.Println(sl3)
	fmt.Println(sl4)

	// sl4にsl3の値をコピー
	n := copy(sl4, sl3)
	fmt.Println(n, sl4) // 5 [1 2 3 4 5]
}