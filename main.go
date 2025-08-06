package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/phrkdll/doomsday/components"
	"github.com/phrkdll/must/pkg/must"
)

const (
	host = "localhost"
	port = "8081"
)

func main() {
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.RFC3339,
		}),
	))

	slog.Info("Doomsday server...")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		components.Home().Render(r.Context(), w)
	})
	mux.HandleFunc("GET /about", func(w http.ResponseWriter, r *http.Request) {
		components.About().Render(r.Context(), w)
	})

	slog.Info("Server is listening...", "host", host, "port", port)

	addr := fmt.Sprintf("%s:%s", host, port)
	must.Succeed(http.ListenAndServe(addr, mux)).ElsePanic()
}
