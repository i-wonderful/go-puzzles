package main

import "fmt"

// Задача: что выведет код
// https://t.me/golangtests/497
func main() {
	a := []int{1, 2, 3}

	var b []int
	copy(b, a)

	c := make([]int, 0, len(a))
	copy(c, a)

	d := make([]int, len(a))
	copy(d, a)

	fmt.Printf("b=%v c=%v f=%v", b, c, d)
}

// Пояснения.
//
// var b []int
// copy(b, a)
// Создается пустой срез b (нулевой длины и емкости).
// Функция copy не скопирует ничего из a в b, так как у b нет выделенной памяти для элементов (длина равна нулю).
// В результате b останется пустым срезом [].
//
// c := make([]int, 0, len(a))
// Создается срез c с нулевой длиной и емкостью, равной длине a (то есть 3).
// Однако copy будет копировать элементы только в существующую длину c, которая равна нулю.
// Поэтому в c ничего не будет скопировано, и он останется пустым срезом [].
//
// d := make([]int, len(a))
// Создается срез d с длиной и емкостью, равной длине a (то есть 3).
// Функция copy скопирует все элементы из a в d, поэтому d будет равен [1, 2, 3].
