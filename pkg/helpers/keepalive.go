package helpers

import (
	"context"
	"net/http"
	"time"
)

// keepalive will send spaces to a web client, until cancelled
func Keepalive(ctx context.Context, w http.ResponseWriter) {
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			_, _ = w.Write([]byte(" ")) // send a space to keep client alive
		}
	}
}
