package main

import "fmt"

// for
func main() {
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("Loop")
	}

	// 条件付きfor
	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	// 3の時スキップ、6で抜ける
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}

	// 配列の繰り返し処理
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// 範囲式
	arr1 := [3]int{1, 2, 3}

	for _, i := range arr1 {
		fmt.Println(i)
	}

	// slice（要素数を指定しない）
	sl := []string{"Python", "PHP", "GO"}
	for i, v := range sl {
		fmt.Println(i, v)
	}
	// 0 Python
	// 1 PHP
	// 2 GO

	// map
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}
	// apple 100
	// banana 200
}