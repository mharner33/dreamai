package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/mharner33/dreamai/handler"
)

//go:embed public
var FS embed.FS

func main() {
	if err := initEnv(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Handle("/", http.FileServer(http.FS(FS)))
	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("Starting server on port:", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initEnv() error {

	return godotenv.Load()
}
