
## Задача. Доступ к репликам бд.
https://youtu.be/x689QxR3AIc?si=sIyZw6ADI2clyP8W&t=48

Есть бд, запущена на тех репликах. Одни и те же данные везде.

Есть функция Get, которая по ключу и реплике получить значение.
Если значение нет, вернется notFound.
Реализовать ф-ю которая конкурентно обращается к этим репликам,
и как только одна из реплик возвращает значение, перестаем ждать ответ 
от других и возвращаем значение.
Т.е. идем параллельно в три реплики, ждем первый результат и возвращаем.

```go
/*
127.0.0.1
127.0.0.2
127.0.0.3
 */

var ErrNotFound = errors.New("not found")

func Get(ctx context.Context, key string) (string, error) {
	return "", nil // already implement
}

func DistributedGet(ctx context.Context, adresses []string, key string) (string, error) {
	// todo implement
}
```