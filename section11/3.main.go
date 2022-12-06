package main

import "fmt"

// structの埋め込み

type T struct {
	User User
}

type User struct {
	Name string
	Age int
}

// ポインタ型
func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t) // {{user1 10}}
	fmt.Println(t.User) // {user1 10}
	fmt.Println(t.User.Name) // user1
	// fmt.Println(t.Name) 直接アクセスする場合は型名省略時のみ

	t.User.SetName()
	fmt.Println(t) // {{A 10}}
}
