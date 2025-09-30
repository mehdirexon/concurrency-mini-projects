package main

import (
	"errors"
	config "final-project/internal/config"
	"final-project/internal/database"
	"final-project/internal/helpers"
	"final-project/internal/middlewares"
	"final-project/internal/render"
	"final-project/internal/router"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	// Set up app config
	var app = &config.AppConfig{
		InfoLogger:   nil,
		ErrorLogger:  nil,
		UseCache:     false,
		InProduction: false,
	}

	// Set up logging
	app.InfoLogger = log.New(os.Stdout, "ℹ️ INFO\t", log.Ldate|log.Ltime)
	app.ErrorLogger = log.New(os.Stdout, "❌ ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Init modules
	database.Register(app)
	render.Register(app)
	helpers.Register(app)
	middlewares.Register(app)

	// Connect to the database
	db := database.Init()
	err = db.Ping()
	if err != nil {
		return
	}

	// Create Session
	app.Session = config.SessionInit()

	// Create channels

	// Create Wait Group
	app.Wait = &sync.WaitGroup{}
	// Set up the application config

	// Set up mail

	// Listen for web connections
	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router.Register(),
	}

	app.InfoLogger.Println("Starting server on port " + os.Getenv("PORT"))

	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		app.ErrorLogger.Panic(err)
	}
}
