package middleware

import "net/http"

// RequireToken (★) пропускает GET всем, а остальные запросы — только с заголовком X-Token, равным token.
// Пока ★ не сделана, это заглушка: оставь её как есть.
func RequireToken(token string, next http.Handler) http.Handler {
	return next
}
