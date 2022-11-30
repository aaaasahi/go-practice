package main

import "fmt"

// init
// initの関数はmainの前に実行される

func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Main")
}