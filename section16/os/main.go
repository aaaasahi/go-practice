package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	// Exit
	// 処理を終了させる
	os.Exit(1) // ここで処理が終了
	fmt.Println("start")

	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(1) // deferが実行されずここで終了


	// log.Fatal
	// エラーを出して処理を終了させる（os.Exitを実行する）
	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
	}


	// Args
	// コマンドラインで与えられたパラメータを受け取る
	// 実行コマンド ./getargs 123 456 789
	fmt.Println(os.Args[0]) // ./getargs コマンド名が入る
	fmt.Println(os.Args[1]) // 123
	fmt.Println(os.Args[2]) // 456
	fmt.Println(os.Args[3]) // 789

	// os.Argsの要素数を表示
	fmt.Printf("length=%d\n", len(os.Args)) // length=4
	// os.Argsの中身を全て表示
	for i, v := range os.Args {
		fmt.Println(i, v) // 0 ./getargs
											// 1 123
											// 2 456
											// 3 789
	}


	// ファイル操作
	// Open
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	// deferでクローズすることで確実にファイルがクローズされる。
	defer f.Close()


	// ファイル操作
	// Create
	// fooファイル作成
	f, _ := os.Create("foo.txt")
	f.Write([]byte("Hello\n"))
	f.WriteAt([]byte("Golang"), 6)
	f.Seek(0, os.SEEK_END)
	f.WriteString("Yaah")


	// ファイル操作
	// Read
	// fooファイルの読み込み
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)

	n,err := f.Read(bs)
	fmt.Println(n)           // 16
	fmt.Println(string(bs))  // Hello GolangYaah

	bs2 := make([]byte, 128)

	nn, err := f.ReadAt(bs2, 10)
	fmt.Println(nn)              // 6
	fmt.Println(string(bs2))     // ngYaah


	// ファイル操作
	// OpenFile // より詳細にファイルをOpenできる
	// O_RDONLY 読み込み専用
	// O_WRONLY 書き込み専用
	// O_RDWR 読み書き可能
	// O_APPEND ファイル末尾に追記
	// O_CREATE ファイルがなければ作成
	// O_TRUNC 可能であればファイルの内容をオープン時に空にする
	f. err := os.Openfile("foo.txt", os.O_RDONLY|O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}
