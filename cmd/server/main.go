package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/Gornak40/standchill/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	defPort = 8080
)

func main() {
	port := flag.Int("port", defPort, "Port for server")
	flag.Parse()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	addr := fmt.Sprintf(":%d", *port)
	slog.Info("init server", slog.String("addr", addr))
	if err := http.ListenAndServe(addr, r); err != nil {
		utils.Fatal("failed to start server", utils.ErrAttr(err))
	}
}
