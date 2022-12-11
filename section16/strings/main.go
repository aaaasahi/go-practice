package main

import (
	"fmt"
	"strings"
)

func main() {
	// 文字列を結合する
	s1 := strings.Join([]string{"A", "B", "C"}, ",")
	s2 := strings.Join([]string{"A", "B", "C"}, "")
	fmt.Println(s1, s2) // A.B.C ABC

	// 文字列に含まれる部分文字列を検索する
	// 検索対象のindex番号を出力
	i1 := strings.Index("ABCDE", "A")
	i2 := strings.Index("ABCDE", "D")
	fmt.Println(i1, i2) // 0 3

	// 最後のindex番号
	i3 := strings.LastIndex("ABCDABCD", "BC")
	fmt.Println(i3) // 5

	// 第二引数の中から最初に現れるindex番号
	i4 := strings.IndexAny("ABC", "ABC")
	// 最後に現れるindex番号
	i5 := strings.LastIndexAny("ABC", "ABC")
	fmt.Println(i4, i5) // 0 2

	// 検索対象が最初に始まるかと最後で終わるか
	b1 := strings.HasPrefix("ABC", "A")
	b2 := strings.HasSuffix("ABC", "C")
	fmt.Println(b1, b2) // true true

	// 検索対象の文字列が含まれているか、複数含まれているか
	b3 := strings.Contains("ABC", "B")
	b4 := strings.ContainsAny("ABCDE", "BD")
	fmt.Println(b3, b4) // true true

	// いくつ含まれているか 空文字の場合は文字列の長さ+1
	i6 := strings.Count("ABCABC", "B")
	i7 := strings.Count("ABCABC", "")
	fmt.Println(i6, i7) // 2 7

	// 文字列を繰り返して結合する
	s3 := strings.Repeat("ABC", 4)
	s4 := strings.Repeat("ABC", 0)
	fmt.Println(s3, s4) // ABCABCABCABC

	// 文字列の置換
	s5 := strings.Replace("AAAAAA", "A", "B", 1)
	s6 := strings.Replace("AAAAAA", "A", "B", -1)
	fmt.Println(s5, s6) // BAAAAA BBBBBB

	// 文字列の分割
	s7 := strings.Split("A,B,C,D,E", ",")
	fmt.Println(s7) // [A B C D E]
	s8 := strings.SplitAfter("A,B,C,D,E", ",")
	fmt.Println(s8) // [A, B, C, D, E]
	s9 := strings.SplitN("A,B,C,D,E", ",", 2)
	fmt.Println(s9) // [A,B,C,D,E]
	s10 := strings.SplitAfterN("A,B,C,D,E", ",", 4)
	fmt.Println(s10) // [A, B, C, D,E]

	// 大文字小文字の変換
	s11 := strings.ToLower("ABC")
	s12 := strings.ToLower("E")
	s13 := strings.ToUpper("abc")
	s14 := strings.ToUpper("e")
	fmt.Println(s11, s12, s13, s14) // abc e ABC E

  // 文字列から空白を取り除く
	s15 := strings.TrimSpace("    -    ABC    -    ") // 前後のスペースを取り除く
	s16 := strings.TrimSpace("　　　　ABC　　　　")
	fmt.Println(s15, s16) // -    ABC    - ABC

	// 文字列からスペースで区切られたフィールドを取り出す
	s18 := strings.Fields("a b c")
	fmt.Println(s18) // [a b c]
	fmt.Println(s18[1]) // b
}
