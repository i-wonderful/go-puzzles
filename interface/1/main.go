package main

import "fmt"

// https://youtu.be/yM7Z6H1G-rE?si=h1wuiqOOLXho_JbJ
type I interface {
	Foo()
}

type S struct {
}

func (s *S) Foo() {
	fmt.Println("foo")
}

func Build() I {
	var res *S
	return res
}

func main() {
	i := Build()
	if i != nil { // тут ничего не сломается, несмотря на то что i не инициализирован
		i.Foo()
	} else {
		fmt.Println("nil")
	}
}
