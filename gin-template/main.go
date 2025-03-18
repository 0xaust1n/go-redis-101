package main

import (
	"log"

	"0xaust1n.github.com/gin-template/internal/pkg/core"
	"0xaust1n.github.com/gin-template/internal/routers"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Create a new HTTP server
	server := core.NewHttpServer()

	// Register routers from the standalone file
	routers.RegisterRouters(server)

	// Run the server
	server.Run(":8080")
}
