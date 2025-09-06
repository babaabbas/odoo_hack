package main

import (
	"fmt"
	"net/http"
	"os"
	"synergy/internal/config"
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
	go func() {
		server.ListenAndServe()
	}()
	<-done

	//storage setup

}
