package config

import (
	"final-project/internal/models"
	"log"
	"sync"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
	Session       *scs.SessionManager
	Wait          *sync.WaitGroup
	MailChan      chan models.Message
	MailErrorChan chan error
	MailDoneChan  chan bool
	UseCache      bool
	InProduction  bool
}
