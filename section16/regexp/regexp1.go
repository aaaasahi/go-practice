package main

import (
	"fmt"
	"regexp"
)

// 正規表現のためのパッケージ
func main() {
	// Goの正規表現の基本 regexp.MatchString
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match) // true

	// Compile
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match) // true

	// MustCompile
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match) // true

	// 正規表現のフラグ
	re3 := regexp.MustCompile(`(?i)abc`)
	match = re3.MatchString("ABC")
	fmt.Println(match) // true

	// 幅を持たない正規表現
	re4 := regexp.MustCompile(`^ABC$`)
	match = re4.MatchString("ABC")
	fmt.Println(match) // true
	match = re4.MatchString("  ABC  ")
	fmt.Println(match) // false

	// 繰り返しを表す正規表現
	re6 := regexp.MustCompile("a+b*")
	fmt.Println(re6.MatchString("ab")) // true
	fmt.Println(re6.MatchString("a"))  // true
	fmt.Println(re6.MatchString("aaaabbbbbbbb")) // true
	fmt.Println(re6.MatchString("b")) // false

	// 正規表現の文字クラス
	re8 := regexp.MustCompile(`[XYZ]`)
	fmt.Println(re8.MatchString("Y")) // true

	re9 := regexp.MustCompile(`^[0-9A-Za-z_]{3}$`)
	fmt.Println(re9.MatchString("ABC")) // true
	fmt.Println(re9.MatchString("abcdefg")) // false

	re10 := regexp.MustCompile(`[^0-9A-Za-z_]`)
	fmt.Println(re10.MatchString("ABC")) // false
	fmt.Println(re10.MatchString("あ"))  // true

  // 正規表現のグループ
	re11 := regexp.MustCompile(`(abc|ABC)(xyz|XYZ)`)
	fmt.Println(re11.MatchString("abcxyz")) // true
	fmt.Println(re11.MatchString("ABCXYZ")) // true
	fmt.Println(re11.MatchString("ABCxyz")) // true
	fmt.Println(re11.MatchString("ABCabc")) // false

	// 正規表現にマッチした文字列の取得
	// FindString
	fs1 := re11.FindString("AAAAABCXYZZZZ")
	fmt.Println(fs1) // ABCXYZ
	fs2 := re11.FindAllString("ABCXYZABCXYZ", -1)
	fmt.Println(fs2) // [ABCXYZ ABCXYZ]

	// 正規表現のグループによるサブマッチ
	// FindAllStringSubmatch
	re12 := regexp.MustCompile(`(\d+)-(\d+)-(\d+)`)
	s := `
	0123-456-7889
	111-222-333
	556-787-899
	`

	ms := re12.FindAllStringSubmatch(s, -1)
	fmt.Println(ms) // [[0123-456-7889 0123 456 7889] [111-222-333 111 222 333] [556-787-899 556 787 899]]

	for _, v := range ms {
		fmt.Println(v) // [0123-456-7889 0123 456 7889]
									 // [111-222-333 111 222 333]
									 // [556-787-899 556 787 899]
	}

	fmt.Println(ms[0][0]) // 0123-456-7889
	fmt.Println(ms[0][1]) // 0123
	fmt.Println(ms[0][2]) // 456
}
