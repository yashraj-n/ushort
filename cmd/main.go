package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/httprate"
	httprateredis "github.com/go-chi/httprate-redis"
	"github.com/joho/godotenv"
	"github.com/yashraj-n/ushort/api"
	"github.com/yashraj-n/ushort/repository"
)

func main() {
	err := godotenv.Load()
	PORT := os.Getenv("PORT")

	if err != nil {
		panic(err)
	}

	if PORT == "" {
		slog.Warn("PORT is not set, using default port 8080")
		PORT = "8080"
	}

	slog.Info("Starting server...")
	router := chi.NewRouter()

	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(httprate.Limit(20, time.Minute*30, httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint), httprateredis.WithRedisLimitCounter(&httprateredis.Config{
		Client: repository.GetRedisInstance(),
	})))

	router.Get("/", api.HandleIndex)
	router.Post("/submit", api.HandleSubmit)
	router.Get("/{hash}", api.HandleRedirect)

	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: router,
	}

	slog.Info("Server started on http://localhost:" + PORT)

	err = server.ListenAndServe()

	if err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
