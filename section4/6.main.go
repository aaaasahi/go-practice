package main

import "fmt"

// 配列型
func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2) // => [A, B ]

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3) // => [1, 2, 3]

	arr4 := [...]string{"C"}
	fmt.Println(arr4) // => [C]
	fmt.Printf("%T\n", arr4) // int[1]

	fmt.Println(arr2[0]) // => [A]
	fmt.Println(arr2[1]) // => [b]
	fmt.Println(arr2[2]) // => [ ]

	arr2[2] = "C"
	fmt.Println(arr2) // => [A, B, C]

	// 要素数出力
	fmt.Println(len(arr1)) // => 3
}