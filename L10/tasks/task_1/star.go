package main

import "errors"

// FromAny собирает товар из значения неизвестного типа (★, часть 1).
// Пока ★ не сделана, это заглушка: оставь её как есть.
func FromAny(raw any) (Product, error) {
	return Product{}, errors.New("не сделано")
}

// Header возвращает названия колонок таблицы через запятую (★, часть 2).
// Пока ★ не сделана, это заглушка: оставь её как есть.
func Header(v any) string {
	return ""
}
