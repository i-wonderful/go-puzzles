package main

import "sync"

// Задание: исправить код
// https://youtu.be/yM7Z6H1G-rE?si=He2Ebg2dE4QoWdeu&t=2671
func main2() {
	m := map[int]int{}

	wg := sync.WaitGroup{}
	var res chan int
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			m[i] = i
			res <- i
			wg.Done()
		}()
	}

	close(res)
	wg.Wait()
}

// Проблемы: канал не инизиализирован (будет паника при закрытии канала)
// пишем в мапу не потокобезопасно
// небезопасное поведение с переменной i

func main() {
	m := map[int]int{}

	wg := sync.WaitGroup{}
	mu := sync.Mutex{} // new

	res := make(chan int, 10) // new
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) { // new
			mu.Lock()         // new
			defer mu.Unlock() // new

			m[i] = i
			res <- i
			wg.Done()
		}(i)
	}

	wg.Wait() // change
	close(res)
}
