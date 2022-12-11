package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 正規表現による文字列の置換
	re1 := regexp.MustCompile(`\s+`)

	fmt.Println(re1.ReplaceAllString("AAA BBB CCC", ",")) // AAA,BBB,CCC

	re2 := regexp.MustCompile(`、|。`)
	fmt.Println(re2.ReplaceAllString("私は、Golangを使用する、プログラマー。", "")) // 私はGolangを使用するプログラマー

	// 正規表現による文字列の分割
	re3 := regexp.MustCompile((`(abc|ABC)(xyz|XYZ)`))
	fmt.Println(re3.Split("ASHVJV<HABCXYZKNJBJVKABCXYZ", -1)) // [ASHVJV<H KNJBJVK ]

	re4 := regexp.MustCompile(`\s+`)
	fmt.Println(re4.Split("aaaaaaaaaa     bbbbbbb  cccccc", -1)) // [aaaaaaaaaa bbbbbbb cccccc]
}
