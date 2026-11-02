// Package catalog — каталог товаров в памяти и его HTTP-обработчики.
package catalog

import (
	"net/http"
	"sync"
)

// Product — товар каталога. Добавь теги json — см. task.md.
type Product struct {
	ID      int
	Name    string
	Price   int
	InStock bool
}

// Validate проверяет название и цену (проверки id здесь нет — см. task.md).
func (p Product) Validate() error {
	return nil // TODO
}

// Store — in-memory хранилище товаров. Поля уже готовы, их менять не нужно:
// запросы идут в разных goroutines, поэтому всё, что лежит в Store, защищает mu.
type Store struct {
	mu     sync.Mutex      // закрывай им любое обращение к items и nextID: mu.Lock() … mu.Unlock()
	items  map[int]Product // товары по номеру
	nextID int             // номер, который получит следующий добавленный товар
}

// NewStore создаёт пустой каталог. Эта функция готова.
// map создан заранее: запись в нулевой (nil) map вызывает panic.
// nextID начинается с 1: по условию первый товар получает id 1, а нулевое значение int — это 0.
func NewStore() *Store {
	return &Store{items: map[int]Product{}, nextID: 1}
}

// Handler возвращает маршруты каталога.
func (s *Store) Handler() http.Handler {
	return http.NewServeMux() // TODO
}
