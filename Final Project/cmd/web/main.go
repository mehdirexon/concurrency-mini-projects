package main

import (
	"encoding/gob"
	"errors"
	"final-project/internal/config"
	"final-project/internal/helpers"
	"final-project/internal/http/handlers"
	"final-project/internal/http/middlewares"
	"final-project/internal/http/render"
	router2 "final-project/internal/http/router"
	"final-project/internal/models"
	"final-project/internal/service/mail"
	"final-project/internal/store"
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
	store.Register(app)
	render.Register(app)
	helpers.Register(app)
	middlewares.Register(app)

	// Repo
	router2.New(handlers.GetRepo())
	mailer := mail.New(app)

	// Connect to the database
	db := store.Init()
	err = db.Ping()
	if err != nil {
		return
	}
	handlers.Register(app, store.New(db))

	// Create Session
	gob.Register(store.User{})
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Store = redisstore.New(store.RedisInit())
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure, _ = strconv.ParseBool(os.Getenv("PRODUCTION"))
	app.Session = session

	// Create channels
	mailerChan := make(chan models.Message, 100)
	mailerDoneChan := make(chan bool)
	mailErrorChan := make(chan error)

	errorChan := make(chan error)
	errorDoneChan := make(chan bool)

	app.ErrorChan = errorChan
	app.ErrorDoneChan = errorDoneChan

	// Create Wait Group
	app.Wait = &sync.WaitGroup{}

	// Set up mailer
	app.MailChan = mailerChan
	app.MailErrorChan = mailErrorChan
	app.MailDoneChan = mailerDoneChan

	// Run the background go routines
	go mailer.ListenForMail()

	// Listen for signal
	go listenForShutdown()

	// Listen for errors
	go listenForErrors()

	// Listen for web connections
	srv := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router2.Register(),
	}

	app.InfoLogger.Println("Starting server on port " + os.Getenv("PORT"))

	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		app.ErrorLogger.Panic(err)
	}
}
func listenForErrors() {
	for {
		select {
		case err := <-app.ErrorChan:
			app.ErrorLogger.Println(err)
		case <-app.ErrorDoneChan:
			return
		}
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
	app.ErrorDoneChan <- true

	app.InfoLogger.Println("Closing channels...")

	close(app.MailChan)
	close(app.MailDoneChan)
	close(app.MailErrorChan)

	close(app.ErrorDoneChan)
	close(app.ErrorChan)
}
