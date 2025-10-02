package main

import (
	"encoding/gob"
	"errors"
	"final-project/internal/config"
	"final-project/internal/database"
	"final-project/internal/handlers"
	"final-project/internal/helpers"
	"final-project/internal/mail"
	"final-project/internal/middlewares"
	"final-project/internal/models"
	"final-project/internal/render"
	"final-project/internal/router"
	_ "final-project/internal/store"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
)

// Set up app config
var app = &config.AppConfig{
	InfoLogger:   nil,
	ErrorLogger:  nil,
	UseCache:     false,
	InProduction: false,
}

func main() {
	// Load environment variables
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}

	// Set up logging
	app.InfoLogger = log.New(os.Stdout, "ℹ️ INFO\t", log.Ldate|log.Ltime)
	app.ErrorLogger = log.New(os.Stdout, "❌ ERROR\t", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	// Init modules
	database.Register(app)
	render.Register(app)
	helpers.Register(app)
	middlewares.Register(app)

	router.New(handlers.GetRepo())
	mailer := mail.New(app)

	// Repo stuffs

	// Connect to the database
	db := database.Init()
	err = db.Ping()
	if err != nil {
		return
	}
	handlers.Register(app, database.New(db))

	// Create Session
	gob.Register(database.User{})
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Store = redisstore.New(config.RedisInit())
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))
	app.Session = session

	// Create channels
	errorChan := make(chan error)
	mailerChan := make(chan models.Message, 100)
	mailerDoneChan := make(chan bool)

	// Create Wait Group
	app.Wait = &sync.WaitGroup{}

	// Create a model

	// Set up the application config

	// Set up mailer
	app.MailChan = mailerChan
	app.MailErrorChan = errorChan
	app.MailDoneChan = mailerDoneChan

	go mailer.ListenForMail()

	// Listen for signal
	go listenForShutdown()

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

func listenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	shutdown()
	os.Exit(0)
}

func shutdown() {
	// cleaning up
	app.InfoLogger.Println("Cleaning up...")

	app.Wait.Wait()

	app.MailDoneChan <- true

	app.InfoLogger.Println("Closing channels...")

	close(app.MailChan)
	close(app.MailDoneChan)
	close(app.MailErrorChan)
}
