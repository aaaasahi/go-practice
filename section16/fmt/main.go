package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("表示")

	// 標準
	fmt.Print("Hello")
	// 改行
	fmt.Println("Hello!")

	// 連結して表示
	fmt.Println("Hello", "world!")
	fmt.Println("Hello", "world!")

	// フォーマットを指定 型を調べるときに使用
	fmt.Printf("%s\n", "Hello")
	fmt.Printf("%#v\n", "Hello")

	// 文字列を生成
	fmt.Sprint("Hello")
	fmt.Sprintf("Hello")
	fmt.Sprintln("Hello")

	// 書き込み先を指定
	fmt.Fprint(os.Stdout, "Hello")
	fmt.Fprintf(os.Stdout, "Hello")
	fmt.Fprintln(os.Stdout, "Hello")

	f, _ := os.Create("test1.txt")
	defer f.Close()

	fmt.Println(f, "Fprint") // test1.txtにFprintを書き込み
}
