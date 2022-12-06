package main

import "fmt"

// structメソッド

type User struct {
	Name string
	Age int
}

// structを使った関数
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

// ポインタ型で宣言 基本はポインタ型で定義する
func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName() // user1

	user1.SetName("A")
	user1.SayName() // user1 更新できていない

	user1.SetName2("A")
	user1.SayName() // A 引数にポインタを使うことで更新ができる

	user2 := &User{Name: "user2"}
	user2.SetName2("B")
	user2.SayName() // B
}