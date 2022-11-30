package main

import "fmt"

// switch
func main() {
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("not")
	}

	switch n := 3; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("not")
	}

	// 条件式パターン
	n1 := 6
	switch {
	case n1 > 0 && n1 < 4:
		fmt.Println("0 < n < 4")
	case n1 > 3 && n1 < 7:
		fmt.Println("3 < n < 7")
	}
}