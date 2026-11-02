// Package logging создаёт логгер для приложений магазина.
package logging

import (
	"io"
	"log/slog"
)

// New возвращает логгер, пишущий в w. Формат "json" — JSON, любой другой — обычный текст.
func New(w io.Writer, format string) *slog.Logger {
	return slog.Default() // TODO
}
