package main

// Что выведет код?
// https://tech-questions.notion.site/Default-value-in-map-5a149debb69348f0838d3c289ba0f065
func main() {
	m := make(map[string]int, 3)
	x := len(m)
	m["Go"] = m["Go"]
	y := len(m)
	println(x, y)
}

// Пояснение. Строка 8 m["Go"] = m["Go"]
// Это выражение выглядит как добавление значения по ключу "Go". Однако, если ключ "Go" отсутствует в карте, то m["Go"] вернёт нулевое значение (в случае int это 0).
//Таким образом, фактически происходит добавление пары "Go": 0 в карту.
