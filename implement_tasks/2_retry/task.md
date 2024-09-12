
Задача из прошлой части.

Сделать ретрай если не прошли в какую то реплику.
Делать ретрай если нет значений в канале результата.

```go
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
			val, err := Get(ctxWithTimeout, address, key)
			 
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
```