package main

import (
	"fmt"
)

// Задача: что выведет код?
// https://tech-questions.notion.site/Append-to-slice-028daf6c644b43bfb58a0580334f8bfa
func main() {
	fmt.Println(arr())
}

func arr() []int {
	arr := []int{0, 1, 2, 3, 4}
	arr[0] = 1
	arr = append(arr, 5)
	newArr := append(arr, 6)
	arr[0] = 10
	return newArr
}

// Пояснение. на строчке arr = append(arr, 5) происходит перерастределение памяти,
// и теперь размер внутреннего массива arr cap=10,
// с следующих строках произойдет изменение внутреннего массива и все изменения отобразятся.
