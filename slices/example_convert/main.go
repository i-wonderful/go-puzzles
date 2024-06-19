package main

import (
	"fmt"
	"sort"
)

// ---------------------------------
// Пример преобразования массива в слайс
// ---------------------------------
// https://habr.com/ru/articles/658623/

func main() {
	var arr = [...]struct{ Name string }{{Name: "b"}, {Name: "c"}, {Name: "a"}}
	//             ^^^^^^^^^^^^^^^^^^^^^ анонимная структура с нужным нам полем

	fmt.Println(arr) // [{b} {c} {a}]

	sort.SliceStable(arr[:], func(i, j int) bool { return arr[i].Name < arr[j].Name })
	//                  ^^^ вот тут вся "магия" - из массива сделали слайс

	fmt.Println(arr) // [{a} {b} {c}]
}
