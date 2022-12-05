package main

import "fmt"

// channel
// 複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造
// 宣言、操作
func main() {

	var ch1 chan int
	// 受信用
	//var ch2 <-chan int

	// 送信用
	//var ch3 chan<- int

	fmt.Println(cap(ch1))

	// makeを使うことで書き込みと読み込みができるようになる
	ch1 = make(chan int)
	ch2 := make(chan int)
	fmt.Println(ch1) // 0
	fmt.Println(ch2) // 0

	// バッファサイズを指定
	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3)) // 5

	// 送信
	ch3 <- 1
	fmt.Println(len(ch3)) // 1

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3)) // 3

	// 受信 キューの性質（先に入れた値から取り出される）
	i := <- ch3
	fmt.Println(i) // 1
	fmt.Println("len", len(ch3)) // 2

	i2 := <- ch3
	fmt.Println(i2) // 2
	fmt.Println("len", len(ch3)) // 1

	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3)) // 0
}
