package main

import (
	"fmt"
	"sort"
)

type Entry struct {
	Name  string
	Value int
}

func main() {

	i := []int{5, 3, 2, 4, 5, 6, 4, 8, 9, 8, 7, 10}
	s := []string{"a", "z", "j"}

	fmt.Println(i, s) // [5 3 2 4 5 6 4 8 9 8 7 10] [a z j]

	// sort
	sort.Ints(i)
	sort.Strings(s)

	fmt.Println(i, s) // [2 3 4 4 5 5 6 7 8 8 9 10] [a j z]

	el := []Entry{
		{"A", 20},
		{"F", 40},
		{"i", 30},
		{"b", 10},
		{"t", 15},
		{"y", 30},
		{"c", 30},
		{"w", 30},
	}

	fmt.Println(el) // [{A 20} {F 40} {i 30} {b 10} {t 15} {y 30} {c 30} {w 30}]

	// Slice
	sort.Slice(el, func(i, j int) bool { return el[i].Name < el[j].Name })

	fmt.Println("---------")
	fmt.Println(el) // [{A 20} {F 40} {b 10} {c 30} {i 30} {t 15} {w 30} {y 30}]
	fmt.Println("---------")

	// SliceStable
	sort.SliceStable(el, func(i, j int) bool { return el[i].Name < el[j].Name })

	fmt.Println("---------")
	fmt.Println(el) // [{A 20} {F 40} {b 10} {c 30} {i 30} {t 15} {w 30} {y 30}]
	fmt.Println("---------")
}
