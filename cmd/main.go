package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/1shoukr/swiftplay-backend/internal/server"
)

func main() {
	srv, err := server.NewServer()
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	// Graceful shutdown
	go func() {
		fmt.Println("Server starting on :8081")
		if err := srv.Run(":8081"); err != nil {
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
