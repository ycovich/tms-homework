// L7 / task_5 — Предскажи вывод.
// НЕ ЗАПУСКАЙ, пока не запишешь свои ответы в answers.txt!
// Каждый фрагмент печатает ровно одну строку вида «N: ...».
package main

import (
	"fmt"
	"slices"
)

func double(x int) {
	x = x * 2
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func addItem(s []int) {
	s = append(s, 4)
	s[0] = 100
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	// 1. Параметр — копия
	{
		x := 5
		double(x)
		fmt.Println("1:", x)
	}

	// 2. Параметры — адреса
	{
		a, b := 1, 2
		swap(&a, &b)
		fmt.Println("2:", a, b)
	}

	// 3. Массив
	{
		a := [3]int{1, 2, 3}
		b := a
		b[0] = 100
		fmt.Println("3:", a[0], b[0])
	}

	// 4. Слайс
	{
		s := []int{1, 2, 3}
		t := s
		t[0] = 100
		fmt.Println("4:", s[0], t[0])
	}

	// 5. Срез
	{
		s := []int{1, 2, 3, 4, 5}
		part := s[1:3]
		part[1] = 0
		fmt.Println("5:", s, len(part), cap(part))
	}

	// 6. append, когда место есть
	{
		a := []int{1, 2, 3, 4}
		b := a[:2]
		b = append(b, 99)
		fmt.Println("6:", a)
	}

	// 7. Срез с третьим числом
	{
		a := []int{1, 2, 3, 4}
		b := a[:2:2]
		b = append(b, 99)
		fmt.Println("7:", a, b)
	}

	// 8. append внутри функции: ёмкость с запасом
	{
		nums := make([]int, 3, 10)
		addItem(nums)
		fmt.Println("8:", nums, len(nums), nums[:4])
	}

	// 9. append внутри функции: запаса нет
	{
		nums := []int{1, 2, 3}
		addItem(nums)
		fmt.Println("9:", nums)
	}

	// 10. range
	{
		nums := []int{1, 2, 3}
		for _, v := range nums {
			v = v * 10
		}
		fmt.Println("10:", nums)
	}

	// 11. Указатель на элемент и append
	{
		s := []int{1, 2, 3}
		p := &s[0]
		s = append(s, 4)
		*p = 100
		fmt.Println("11:", s[0], *p)
	}

	// 12. Clone слайса слайсов
	{
		g := [][]int{{1, 2}, {3, 4}}
		g2 := slices.Clone(g)
		g2[0][0] = 9
		g2[1] = []int{7, 7}
		fmt.Println("12:", g)
	}

	// 13. Замыкания
	{
		c1 := counter()
		c1()
		c1()
		c2 := counter()
		fmt.Println("13:", c1(), c2())
	}
}
