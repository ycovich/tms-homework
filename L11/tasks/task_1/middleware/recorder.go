package middleware

import "net/http"

// statusRecorder запоминает код ответа, который отправил обработчик.
// Сам http.ResponseWriter код не хранит, а логу он нужен. Этот файл готов, менять его не нужно.
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (s *statusRecorder) WriteHeader(code int) {
	s.status = code
	s.ResponseWriter.WriteHeader(code)
}
