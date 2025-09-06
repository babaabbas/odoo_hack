package main

import (
	"fmt"
	"synergy/internal/config"
)

func main() {
	//config
	var cfg config.Config
	cfg = *config.Must_Load()
	fmt.Print(cfg)
	//setup routers
	//setup server
	//storage setup

}
