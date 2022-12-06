package main

import "fmt"

// struct

type User struct {
	Name string
	Age int
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

// 引数をポインタ型
func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1) // { 0} stringとintの初期値
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1) // {user1 10}

	user2 := User{}
	fmt.Println(user2) // { 0}
	user2.Name = "user2"
	fmt.Println(user2) // {user2 0}

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3) // {user3 30}

	user4 := User{"user4", 40}
	fmt.Println(user4) // {user4 40}

	// 順番が違う場合エラー
	// user5 := User{30, "user5"}
	// fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6) // {user6 0}

	// ポインタ型
	user7 := new(User)
	fmt.Println(user7) // &{ 0}

	// ポインタ型
	user8 := &User{}
	fmt.Println(user8) // &{ 0}

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1) // {user1 10} 更新されていない
	fmt.Println(user8) // &{A 1000} ポインタ型で渡したので更新できている
}
