// Package server запускает http.Handler и аккуратно останавливает его. Он ничего не знает о каталоге.
package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
)

// Run слушает addr и обслуживает h, пока не отменён ctx. После отмены перестаёт принимать
// соединения и ждёт выполняющиеся запросы не дольше 10 секунд. При штатной остановке возвращает nil.
func Run(ctx context.Context, addr string, h http.Handler, log *slog.Logger) error {
	return errors.New("не сделано") // TODO
}
