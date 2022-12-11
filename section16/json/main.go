package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

// 構造体からJSONテキストへの変換
type A struct{}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A",omitempty`
}

// Unmarshalのカスタム
func (u *User) UnmarshalJSON(b []byte) error {
	type User2 struct {
		Name string
	}
	var u2 User2
	err := json.Unmarshal(b, &u2)
	if err != nil {
		fmt.Println(err)
	}
	u.Name = u2.Name + "!"
	return err
}

// Marshalのカスタム
func (u User) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + u.Name,
	})
	return v, err
}

func main() {
	u := new(User)
	u.ID = 1
	u.Name = "test"
	u.Email = "example@example.com"
	u.Created = time.Now()

	// Marshal jsonに変換
	bs, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs)) // {"id":1,"name":"test","email":"example@example.com","created":"2022-12-11T21:40:27.39474+09:00","A":null}

	// ----------------------

	fmt.Printf("%T\n", bs)

	u2 := new(User)

	// Unmarshal JSONをデータに変換
	if err := json.Unmarshal(bs, u2); err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("u2", u2)
}
