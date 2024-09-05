package main

import "fmt"

// Задача. Создать множество через мапу и пустую структуру.
// Примечание. Пустая структура struct{} занимает 0 байт памяти и может быть использована маркер события и или состояния

// Создаем тип для представления множества
type Set map[interface{}]struct{}

// Функция для добавления элемента в множество
func (set Set) Add(value interface{}) {
	set[value] = struct{}{}
}

// Функция для удаления элемента из множества
func (set Set) Remove(value interface{}) {
	delete(set, value)
}

// Функция для проверки наличия элемента в множестве
func (set Set) Contains(value interface{}) bool {
	_, exists := set[value]
	return exists
}

func main() {
	// Создаем новое множество
	mySet := make(Set)

	// Добавляем элементы в множество
	mySet.Add("apple")
	mySet.Add("banana")
	mySet.Add("orange")

	// Проверяем наличие элемента в множестве
	fmt.Println("Is 'apple' in the set?", mySet.Contains("apple")) // Вывод: true
	fmt.Println("Is 'grape' in the set?", mySet.Contains("grape")) // Вывод: false

	// Удаляем элемент из множества
	mySet.Remove("banana")

	// Проверяем наличие элемента после удаления
	fmt.Println("Is 'banana' in the set?", mySet.Contains("banana")) // Вывод: false
}
