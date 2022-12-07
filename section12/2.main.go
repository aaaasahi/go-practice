package main

import "fmt"

// interface
// カスタムエラー

type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生しました", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error()) // カスタムエラーが発生しました

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode) // 1234
	}
}