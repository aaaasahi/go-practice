package main

import "fmt"

// ラベル付きfor
// ラベル内のループ処理を抜ける

func main() {
	Loop:
		for {
			for {
				for {
					fmt.Println("START")
					// Loop内の処理を一気に飛ばす
					break Loop
				}
				fmt.Println("処理をしない")
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("END")

		// continueとの併用
	Loop:
		for i := 0; i < 3; i++ {
			for j := 1; j < 3; j++ {
				if j > 1 {
					continue Loop
				}
				fmt.Println(i, j, i*j)
			}
			fmt.Println("処理をしない")
		}
}