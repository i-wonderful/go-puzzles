package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
127.0.0.1
127.0.0.2
127.0.0.3
*/

var ErrNotFound = errors.New("not found")
var ErrDistributedGetFailed = errors.New(" err ")

func Get(ctx context.Context, address, key string) (string, error) {
	return "", nil // already implement
}

// тут реализация
func DistributedGet(ctx context.Context, adresses []string, key string) (string, error) {
	if ctx.Err() != nil {
		return "", ctx.Err()
	}

	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(len(adresses))

	valueCh := make(chan string)
	notFoundErrCh := make(chan error)

	go func() {
		wg.Wait()
		close(valueCh)
		close(notFoundErrCh)
	}()

	for _, address := range adresses {
		go func(address string) {
			defer wg.Done()

			var val string
			var err error

			withRetry(func() (error, string) {
				val, err = Get(ctxWithTimeout, address, key)
				return err, val
			}, 3, 1)

			if errors.Is(err, ErrNotFound) {
				select {
				case notFoundErrCh <- err:
				default:
				}
				return
			} else if err != nil {
				return
			}

			select {
			case valueCh <- val:
			default:
			}

		}(address)
	}

	select {
	case <-ctxWithTimeout.Done():
		return "", ctxWithTimeout.Err()
	case res, ok := <-valueCh:
		if ok == false {
			return "", ErrDistributedGetFailed
		}
		return res, nil
	case resErr, ok := <-notFoundErrCh:
		if ok == false {
			return "", ErrDistributedGetFailed
		}
		return "", resErr
	}
}

// функция которая обращается с ретраем
func withRetry(cb func() (error, string), count int, delay time.Duration) (string, error) {

	for count >= 0 {
		count++
		err, val := cb()

		if errors.Is(err, ErrNotFound) {
			return "", err
		}

		if err != nil {
			continue
		}

		return val, nil
		time.Sleep(delay)
		delay *= 2
	}

	// todo ?
	return "", fmt.Errorf("some err")
}

func main() {

}
