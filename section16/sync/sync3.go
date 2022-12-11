package main

import (
	"fmt"
)

// ゴルーチンの終了を待ち受ける

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
		}
	}()

	// 無限ループをつけないと上記が実行されない
	/*
		for {

		}
	*/
}
