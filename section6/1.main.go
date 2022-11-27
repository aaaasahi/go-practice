package main

import "fmt"

// 算術演算子
func main() {
	fmt.Println(1 + 2)
	fmt.Println("ABC" + "DE")
	fmt.Println(5 - 1)
	fmt.Println(5 * 4)
	fmt.Println(60 / 3)
	fmt.Println(9 % 4)

	n := 0
	n += 4
	fmt.Println(n) // n = n + 4
	n++
	fmt.Println(n) // n = n + 1
	n--
	fmt.Println(n) // n = n - 1

	s := "ABC"
	s += "DEF"
	fmt.Println(s) // "ABC" + "DEF"
}