package main

import (
	"fmt"
)

// Вопрос: что выведет код?
// https://youtu.be/yM7Z6H1G-rE?si=4CHfAzRhP-2jWxkj&t=2540

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	//time.Sleep(time.Second * 3)
	fmt.Println("done")
}

// Ответ: выведет число (или не выведет), и done, т.к горутины не синхронизированы
