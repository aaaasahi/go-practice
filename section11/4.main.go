package main

import "fmt"

// struct コンストラクタ

type User struct {
	Name string
	Age int
}

// コンストラクタ
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("user1", 10)
	fmt.Println(user1) // &{user1 10}
	// 実体にアクセス
	fmt.Println(*user1) // {user1 10}
}