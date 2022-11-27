package main

import "fmt"

// 比較演算子
func main() {
	fmt.Println(1 == 1) // true
	fmt.Println(1 == 2) // false

	fmt.Println(4 <= 8) // true
	fmt.Println(1 >= 8) // false

	fmt.Println(1 < 8)  // true
	fmt.Println(3 > 1)  // true

	fmt.Println(true == false) // false
	fmt.Println(true != false) // true
}