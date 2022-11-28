package main

import (
	"fmt"
	"strconv"
)

// エラーハンドリング
func main() {
	var s string = "100"
	i, err := strconv.Atoi(s)
	// errに値が入っているかどうか捕捉する
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("i = %T\n", i)
}