package main

import (
	"go-user-auth/config"
	"go-user-auth/db/repository"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-user-auth/routers"

	"github.com/gofiber/fiber/v2"
)

// @version 3.0
// @description  Proptuity API
// @BasePath /api
// @in                          header
// @name                        Authentication
// @description					JWT Authentication header for all requests
func main() {
	// app := fiber.New()

	// Load Configurations
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Error loading configurations: %v", err)
	}

	log.Println("Configurations loaded: ", cfg)

	// Todo: Connect to database
	err = repository.Init(cfg)
	if err != nil {
		log.Fatal("unable to connect to storage instances", err.Error())
	}

	// Check if Public and Private keys are stored in the database
	// If not, generate new keys and store them in the database
	KeyId, KeySecret, err := repository.InitKeys(cfg)
	if err != nil {
		log.Fatal("unable to generate keys", err.Error())
	}

	cfg.KeyId = KeyId
	cfg.KeySecret = KeySecret

	// Setup routers
	app := routers.SetupRoutes(cfg)

	// Start server
	go startServer(cfg, app)
	gracefulShutdown(app)
}

// startServer listens to the configured port and starts the server
func startServer(cfg config.AccountsConfig, app *fiber.App) {
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatal("error starting server:", err)
	}
}

// gracefulShutdown handles the shutdown of the server when an interrupt or terminate signal is received
func gracefulShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// attempt to gracefully shut down the server
	if err := app.Shutdown(); err != nil {
		log.Fatal("error shutting down server:", err)
	}
}
