package main

import "fmt"

// if
func main() {
	a := 1
	if a == 2 {
		fmt.Println("two")
	} else if a == 1 {
		fmt.Println("one")
	} else {
		fmt.Println("not")
	}

	// 簡易文付きif
	if b := 100; b == 100 {
		fmt.Println("100")
	}

	// スコープ注意
	x := 0
	if x := 2; true {
		fmt.Println(x) // 2
	}
	fmt.Println(x) // 0
}