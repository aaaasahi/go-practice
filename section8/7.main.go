package main

import (
	"fmt"
	"os"
)

// defer
// 関数が終了した時の処理を登録

func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

// deferの複数登録は後ろから実行される
func RunDefer() {
	defer fmt.Println("RunDefer1")
	defer fmt.Println("RunDefer2")
	defer fmt.Println("RunDefer3") // ここから実行
}

// file作成
func CreateFile() {
	file, err := os.Create("test.text")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	file.Write([]byte("Hello"))
}

func main() {
	TestDefer()

	// 無名関数を使って複数処理の登録
	defer func() {
		fmt.Println("1")
		fmt.Println("2")
		fmt.Println("3")
	}()

	RunDefer()
	CreateFile()
}