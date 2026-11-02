// Package middleware — обёртки вокруг http.Handler. Они не знают ни о каталоге, ни о конкретных маршрутах.
package middleware

import (
	"log/slog"
	"net/http"
)

// RequestID ставит каждому ответу заголовок X-Request-ID: 1, 2, 3…
func RequestID(next http.Handler) http.Handler {
	return next // TODO
}

// Logging пишет в log по одной строке на запрос: method, path, status, duration.
// Для status используй готовый struct statusRecorder из recorder.go.
func Logging(log *slog.Logger, next http.Handler) http.Handler {
	return next // TODO
}
