package main

import (
	"log"

	"github.com/natanaelrusli/segokuning-be/internal/config"
	"github.com/natanaelrusli/segokuning-be/internal/httpserver"
)

func main() {
	log.Println("Welcome to segokuning!")

	cfg := config.InitConfig()

	httpserver.InitGinServer(cfg)
}
