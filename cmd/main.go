package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/1shoukr/swiftplay-backend/internal/server"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	port := ":" + strconv.Itoa(srv.Config.Port)

	// Graceful shutdown
	go func() {
		fmt.Printf("Server starting on %s\n", port)
		if err := srv.Run(port); err != nil {
			log.Fatal("Server failed to start:", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down server...")

	// Close database connection
	if err := srv.Close(); err != nil {
		log.Printf("Error closing server: %v", err)
	}

	fmt.Println("Server stopped")
}
