package main

// Что выведет код?
// https://go101.org/quizzes/scope-2.html

var f = func(x int) {}

func Bar() {
	f := func(x int) {
		if x >= 0 {
			print(x)
			f(x - 1)
		}
	}
	f(2)
}

func Foo() {
	f = func(x int) {
		if x >= 0 {
			print(x)
			f(x - 1)
		}
	}
	f(2)
}

func main() {
	Bar()
	print(" | ")
	Foo()
}
