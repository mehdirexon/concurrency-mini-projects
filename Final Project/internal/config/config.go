package config

import (
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	InfoLogger   *log.Logger
	ErrorLogger  *log.Logger
	Session      *scs.SessionManager
	Wait         *sync.WaitGroup
	UseCache     bool
	InProduction bool
}
