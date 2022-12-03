package main

import "fmt"

// マップ
func main() {
	// マップの作成
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m) // map[A:100 B:200]

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2) // map[A:100 B:200]

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3) // map[1:A 2:B]

	// makeで作成
	m4 := make(map[int]string)
	fmt.Println(m4) // map[]

	// makeで作成したmapに値を追加
	m4[1] = "JAPAN"
	m4[2] = "USA"
	fmt.Println(m4) // map[1:JAPAN 2:USA]

	// 値の取り出し
	fmt.Println(m["A"]) // 100
	fmt.Println(m4[2]) // USA
	fmt.Println(m4[3]) // 値がない場合は初期値

	// 値の取り出しができたかの判定
	s, ok := m4[3]
	fmt.Println(s, ok) // 空とfalse

	// エラーハンドリング
	_, ok1 := m4[3]
	if !ok1 {
		fmt.Println("error") // error
	}

	// 値の更新
	m4[2] = "US"
	fmt.Println(m4) // map[1:JAPAN 2:US]

	m4[3] = "CHAINA"
	fmt.Println(m4) // map[1:JAPAN 2:US 3:CHAINA]

	// 削除
	delete(m4, 3)
	fmt.Println(m4) // map[1:JAPAN 2:US]

	fmt.Println(len(m4)) // 2
}