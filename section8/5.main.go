package main

import "fmt"

// 型アサーションと型switch
// インターフェースの型を動的に判定したいときに使う

func anyting(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 1000)
	}
}

func main() {
	anyting("aaa")
	anyting(1)

	var x interface{} = 3

  // interfaceをintに復元して計算
	i := x.(int)
	fmt.Println(i + 2)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("nil")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString, "x is string")
	} else {
		fmt.Println("none")
	}

	// 型switch
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("none")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	}
}