// L7 / task_14, часть B — найди и исправь ошибку.
// Программа проверяет возраст. Сейчас она считает ошибкой даже правильный возраст.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ValidationError — ошибка проверки: какое поле и что с ним не так
type ValidationError struct {
	Field  string
	Reason string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Reason
}

func validate(age int) error {
	var verr *ValidationError
	if age < 0 {
		verr = &ValidationError{Field: "возраст", Reason: "не может быть отрицательным"}
	} else if age > 150 {
		verr = &ValidationError{Field: "возраст", Reason: "больше 150"}
	}
	return verr
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	bad := 0
	for in.Scan() {
		age, _ := strconv.Atoi(in.Text())
		err := validate(age)
		if err != nil {
			bad++
			fmt.Printf("%d: ошибка: %v\n", age, err)
			continue
		}
		fmt.Printf("%d: ok\n", age)
	}
	fmt.Println("ошибок:", bad)
}
