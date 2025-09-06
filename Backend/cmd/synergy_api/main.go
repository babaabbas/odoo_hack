package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"synergy/internal/config"
	"syscall"
	"time"
)

func main() {
	//config
	cfg := *config.Must_Load()

	fmt.Print(cfg)
	//setup routers
	router := http.NewServeMux()
	router.HandleFunc("GET /api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"List of students"}`))
	})

	//setup server
	server := http.Server{
		Addr:    cfg.Http_Server.Addr,
		Handler: router,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("server is not starting")
		}
	}()
	<-done
	slog.Info("server is shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Info("server not closing ")
	}

	//storage setup

}
