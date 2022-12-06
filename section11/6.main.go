package main

import "fmt"

// struct マップ

type User struct {
	Name string
	Age int
}

func main() {
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}
	fmt.Println(m) // map[1:{user1 10} 2:{user2 20}]

	// keyにUserを指定
	m2 := map[User]string{
		{Name: "user1", Age: 10}: "Tokyo",
		{Name: "user2", Age: 20}: "LA",
	}
	fmt.Println(m2) // map[{user1 10}:Tokyo {user2 20}:LA]

	// makeで作る
	m3 := make(map[int]User)
	fmt.Println(m3) // map[]
	m3[1] = User{Name: "user3"}
	fmt.Println(m3) // map[1:{user3 0}]

	for _, v := range m {
		fmt.Println(v) // {user1 10}
									 // {user2 20}
	}
}
