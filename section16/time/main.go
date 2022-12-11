package main

import (
	"fmt"
	// "log"
	"time"
)

// time

func main() {
	// 時間の生成
	// 今の時間
	t := time.Now()
	fmt.Println(t) // 2022-12-11 14:19:57.34049 +0900 JST m=+0.000221751

	// 指定した時間を生成
	t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)           	 // 2020-06-10 10:10:10 +0900 JST
	fmt.Println(t2.Year())    	 // 2020
	fmt.Println(t2.YearDay()) 	 // 162
	fmt.Println(t2.Month())   	 // June
	fmt.Println(t2.Weekday()) 	 // Wednesday
	fmt.Println(t2.Day())     	 // 10
	fmt.Println(t2.Hour())    	 // 10
	fmt.Println(t2.Minute())  	 // 10
	fmt.Println(t2.Second())  	 // 10
	fmt.Println(t2.Nanosecond()) // 0
	fmt.Println(t2.Zone())       // JST 32400

  // 時間の感覚を表現
	// time.Duration型を返す
	fmt.Println(time.Hour)        // 1h0m0s
	fmt.Printf("%T\n", time.Hour) // time.Duration
	fmt.Println(time.Minute)      // 1m0s
	fmt.Println(time.Second)      // 1s
	fmt.Println(time.Millisecond) // 1ms
	fmt.Println(time.Microsecond) // 1µs
	fmt.Println(time.Nanosecond)  // 1ns

	// time.Duration型を文字列から生成する
	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d) // 2h30m0s

	t3 := time.Now()
	t3 = t3.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t3) // 2022-12-11 14:22:12.340971 +0900 JST m=+135.000703168
	// 現在時刻の2分30秒後を表すtime.Time型の生成


	// 時刻の差分を取得
	t5 := time.Date(2020, 7, 24, 0, 0, 0, 0, time.Local)
	t6 := time.Now()
	fmt.Println(t6) // 2022-12-11 14:29:23.672741 +0900 JST m=+0.000586292

	// t5 - t6
	d2 := t5.Sub(t6)
	fmt.Println(d2) // -20894h29m23.672741s

	// 時刻を比較する 後か前か
	fmt.Println(t6.Before(t5)) // false
	fmt.Println(t6.After(t5))  // true
	fmt.Println(t5.Before(t6)) // true
	fmt.Println(t5.After(t6))  // false

	// 指定時間のスリープ
	// 5秒間停止させて実行
	time.Sleep(5 * time.Second)
	fmt.Println("5秒停止後表示")
}