package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Lionel-Wilson/arritech-takehome/internal/config"
	"github.com/Lionel-Wilson/arritech-takehome/internal/http/router"
	internallogger "github.com/Lionel-Wilson/arritech-takehome/internal/logger"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/storage"
	"github.com/Lionel-Wilson/arritech-takehome/internal/user/storage/entity"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // <-- Add this line to register the Postgres driver
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load environment variables from .env file (only in development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Shutdown on Ctrl+C
	go listenForShutdown(cancel)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	logger := internallogger.New(cfg)

	// Connect with GORM (Postgres)
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		logger.Fatalf("failed to connect to database: %v", err)
	}

	// Auto-migrate schema
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		logger.Fatalf("failed to run migrations: %v", err)
	}

	userRepo := storage.NewUserRepository(db)
	userService := user.NewUserService(logger, userRepo)

	mux := router.New(
		logger,
		userService,
	)

	// Start server with context
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: mux,
	}

	go func() {
		logger.Infof("Server starting on port %s", cfg.Port)

		err = server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatalf("ListenAndServe error: %v", err)
		}
	}()

	// Wait for cancel (SIGINT / SIGTERM)
	<-ctx.Done()

	logger.Infof("Shutting down gracefully...")

	err = server.Shutdown(context.Background())
	if err != nil {
		logger.Fatalf("Server forced to shutdown: %v", err)
	}
}

func listenForShutdown(cancel context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	fmt.Println("Shutdown signal received")
	cancel()
}
