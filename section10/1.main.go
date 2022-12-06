package main

import "fmt"

// ポインタ

func Double(i int) {
  i = i * 2
}

// ポインタ型の引数を取る関数
func Doublev2(i *int) {
  *i = *i * 2
}

func Doublev3(s []int) {
  for i, v := range s {
    s[i] = v * 2
  }
}

func main() {
  var n int = 100
  fmt.Println(n)

  // アドレス
  fmt.Println(&n) // 0x1400001a0a0

  Double(n)

  fmt.Println(n) // 100 関数で二倍してもnは100のまま

  // ポインタを宣言
  var p *int = &n
  fmt.Println(p) // 0x1400001a0a0
  fmt.Println(*p) // 100

  // nの値を置き換えられる
  *p = 300
  fmt.Println(n) // 300
  n = 200
  fmt.Println(*p) // 200

  // 値が置き換わる
  Doublev2(&n)
  fmt.Println(n) // 400

  Doublev2(p)
  fmt.Println(*p) // 800

  // スライスなどの参照型はデフォルトでポインタの機能がある
  var sl []int = []int{1, 2, 3}
  Doublev3(sl)
  fmt.Println(sl) // [2 4 6] ポインタにしなくても値が更新される
}
