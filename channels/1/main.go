package main

import (
	"fmt"
	"time"
)

// ---------------------------------
// https://www.youtube.com/live/oDRqCXw3XGY?si=RUxVPk_3OrcltjuR&t=5792
func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		time.Sleep(time.Second)
		close(ch)
	}()

	for {
		select {
		case v, ok := <-ch:
			if !ok { // без этой проверки будет бесконечно печатать 0 после закрытия канала
				break
			}
			fmt.Println(v)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
			break
		}

	}
}
