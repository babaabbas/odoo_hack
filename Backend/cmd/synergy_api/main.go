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
	"synergy/internal/types"
	"synergy/internal/utils/responses"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
)

func main() {
	fmt.Println(responses.GeneralError(http.ErrAbortHandler))
	//config
	cfg := *config.Must_Load()

	fmt.Print(cfg)
	//setup routers
	router := http.NewServeMux()
	router.HandleFunc("GET /api/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		responses.WriteJson(w, 200, responses.GeneralError(http.ErrBodyNotAllowed))
	})
	var user types.User
	user = types.User{
		ID:           "550e8400-e29b-41d4-a716-446655440000",
		Name:         "avs2",
		Email:        "aliceexample.com",
		PasswordHash: "$2a$10$7sdf8sdf8sdf8sdf8sdf8sdf8sdf8sdf8sdf8sdf8sdf",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := validator.New().Struct(user)
	fmt.Println(responses.ValidateError(err.(validator.ValidationErrors)))

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
