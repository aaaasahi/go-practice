package main

import "fmt"

// panic runtime errorを発生させて強制的に処理を終了させる
// recover panicのruntime errorから復帰できる
// 基本deferと併用
// あまり使用されない

func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("Start")
}