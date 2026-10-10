// L7 / task_14, часть A — Предскажи вывод.
// НЕ ЗАПУСКАЙ, пока не запишешь свои ответы в answers.txt!
// Каждый фрагмент печатает ровно одну строку вида «N: ...».
package main

import "fmt"

type MyErr struct {
	Msg string
}

func (e *MyErr) Error() string { return e.Msg }

// find ничего не нашёл, но вернул «пустой» указатель
func find() error {
	var e *MyErr
	return e
}

type List struct {
	items []int
}

// Len можно вызвать даже у nil-указателя
func (l *List) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

type Shape interface {
	Area() int
}

type Rect struct {
	W, H int
}

func (r Rect) Area() int { return r.W * r.H }

func kind(v any) string {
	switch v.(type) {
	case nil:
		return "nil"
	case error:
		return "error"
	case int:
		return "int"
	default:
		return "другое"
	}
}

func main() {
	// 1. Пустой интерфейс error
	{
		var err error
		fmt.Println("1:", err == nil)
	}

	// 2. nil-указатель внутри error
	{
		err := find()
		fmt.Println("2:", err == nil)
	}

	// 3. Что лежит внутри
	{
		err := find()
		fmt.Printf("3: %T\n", err)
	}

	// 4. %T у пустого интерфейса
	{
		var s fmt.Stringer
		fmt.Printf("4: %T\n", s)
	}

	// 5. Метод у nil-указателя
	{
		var l *List
		fmt.Println("5:", l.Len())
	}

	// 6. any и nil-слайс
	{
		var x any
		a := x == nil
		x = []int(nil)
		b := x == nil
		fmt.Println("6:", a, b)
	}

	// 7. Указатель и any
	{
		var p *int
		var y any = p
		fmt.Println("7:", p == nil, y == nil)
	}

	// 8. Утверждение типа с ok: не угадали
	{
		var v any = 42
		s, ok := v.(string)
		fmt.Printf("8: %q %v\n", s, ok)
	}

	// 9. Утверждение типа с ok: угадали
	{
		var v any = 42
		n, ok := v.(int)
		fmt.Println("9:", n+1, ok)
	}

	// 10. switch по типу
	{
		fmt.Println("10:", kind(nil), kind(find()), kind(3.5))
	}

	// 11. Структура в интерфейсе — это копия
	{
		r := Rect{W: 2, H: 3}
		var s Shape = r
		r.W = 10
		fmt.Println("11:", s.Area(), r.Area())
	}

	// 12. Указатель в интерфейсе
	{
		r := Rect{W: 2, H: 3}
		var s Shape = &r
		r.W = 10
		fmt.Println("12:", s.Area())
	}

	// 13. Структура в map — тоже копия
	{
		m := map[string]Rect{"a": {W: 1, H: 1}}
		v := m["a"]
		v.W = 5
		fmt.Println("13:", m["a"].W, v.W)
	}

	// 14. Сравнение интерфейсов
	{
		var a any = 3
		var b any = 3
		var c any = int64(3)
		fmt.Println("14:", a == b, a == c)
	}
}
