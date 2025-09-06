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
	"synergy/internal/http/handlers/project"
	"synergy/internal/http/handlers/users"
	"synergy/internal/storage/postgres"
	"syscall"
	"time"
)

func main() {
	//config
	cfg := *config.Must_Load()
	pst, err := postgres.New(&cfg)
	if err != nil {
		fmt.Println("Failed to connect to DB:", err)
		return
	}
	fmt.Println("DB object:", pst)
	fmt.Print(cfg)
	//setup routers
	router := http.NewServeMux()
	router.HandleFunc("POST /api/users", users.CreateUserHandler(pst))
	router.HandleFunc("POST /api/projects", project.CreateProjectHandler(pst))
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
	err = server.Shutdown(ctx)
	if err != nil {
		slog.Info("server not closing ")
	}

	//storage setup

}
