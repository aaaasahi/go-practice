package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	// MD5ハッシュ値を生成
	// 任意の文字列からMD5ハッシュ値を生成する処理例
	h := md5.New()
	io.WriteString(h, "ABCDE")

	fmt.Println(h.Sum(nil)) // [46 205 222 57 89 5 29 145 63 97 177 69 121 234 19 109]

	s := fmt.Sprintf("%x", h.Sum(nil))
	fmt.Println(s) // 2ecdde3959051d913f61b14579ea136d
}
