package main

import "fmt"

// interface
// 最もポピュラーなつい買い方。異なる方に共通の性質を付与

// 異なる型をまとめる
type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model:"ABC-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString()) // Name=Taro, Age=21
															// Name=123-456, Age=ABC-1234
	}
}