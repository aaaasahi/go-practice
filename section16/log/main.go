package main

import (
	"log"
	"os"
)

// log

func main() {
	// ログの出力先を標準出力に変更
	log.SetOutput(os.Stdout)

	log.Print("Log\n")
	log.Println("Log2")
	log.Printf("Log%d\n", 3)

	// 止まる
	log.Fatal("Log\n")
	log.Fatalln("Log2")
	log.Fatalf("Log%d\n", 3)

	log.Panic("Log\n")
	log.Panicln("Log2")
	log.Panicf("Log%d\n", 3)

	// ファイルを作成して出力先に指定
	f, err := os.Create("test.log")
	if err != nil {
		return
	}
	log.SetOutput(f)
	log.Println("ファイルに書き込む") // 作成したtest.logに書き込み

	// 標準のログフォーマット
	log.SetFlags(log.LstdFlags)
	log.Println("A")

	// マイクロ秒を追加
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	log.Println("B")

	// 時刻とファイルの行番号（短縮系）
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("C")

	// 標準に戻す
	log.SetFlags(log.LstdFlags)

	// ログのプリフィックスを設定
	log.SetPrefix("[LOG]") // ログの先頭に渡した文字列を出漁
	log.Println("E")

	// ロガーの生成
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Fatalln("message")

	// エラーで終了させる
	_, err = os.Open("fdafdsafa")
	if err != nil {
		// ログ出力
		log.Fatalln("Exit", err)
		// logger.Fatalln("Exit", err)
	}
}
