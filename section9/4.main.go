package main

import "fmt"

// スライス
// for
func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	for i, v := range sl {
		fmt.Println(i, v) // 0 A
											// 1 B
											// 2 C
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i]) // A B C
	}
}