package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// 簡易的なファイルの読み書きができる
func main() {
	// 入力全体を読み込む
	f, _ := os.Open("foo.txt")
	bs, _ := ioutil.ReadAll(f)
	fmt.Println(string(bs1))

	// bar.txtにfoo.txtの内容をファイルに書き込み
	if err := ioutil.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}
}
