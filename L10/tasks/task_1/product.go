package main

// Product — товар в каталоге. Добавь теги к полям — см. task.md.
type Product struct {
	ID      int
	Name    string
	Price   int
	InStock bool
	Note    string
}

// Validate проверяет товар и возвращает первую найденную ошибку.
func (p Product) Validate() error {
	return nil // TODO
}

// FilterValid оставляет только корректные товары. Для каждого отклонённого — ошибка в errs.
func FilterValid(items []Product) (valid []Product, errs []error) {
	return nil, nil // TODO
}

// ParseProducts разбирает JSON-массив товаров.
func ParseProducts(data []byte) ([]Product, error) {
	return nil, nil // TODO
}

// ToJSON превращает товары в JSON-массив.
func ToJSON(items []Product) ([]byte, error) {
	return nil, nil // TODO
}
