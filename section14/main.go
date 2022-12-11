package main

import "fmt"

// テスト

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))
}

// -- テストコマンド --
// テストファイルのあるディレクトリに移動して
// go test
// go test -v 詳細情報
// go test -cover パッケージの関数のテストのカバー率

// ディレクトリ内全てのテストを実行する場合
// go test ./...
