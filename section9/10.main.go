package main

import (
	"fmt"
	"time"
)

// channel
// close

func reciever(name string, ch chan int) {
	for {
		i, ok := <-ch
		if !ok {
			break // closeの場合処理を抜ける
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}
func main() {
	ch1 := make(chan int, 2)

	// チャネルをcloseするとデータの送受信ができなくなる
	// close(ch1)

	// i, ok := <-ch1
	// fmt.Println(i, ok) // 閉じられている場合false

	go reciever("1.goroutin", ch1)
	go reciever("2.goroutin", ch1)
	go reciever("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1) // iが100になったら実行
	time.Sleep(3 * time.Second)
}