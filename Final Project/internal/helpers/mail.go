package helpers

import "final-project/internal/models"

func SendEmail(msg models.Message) {
	app.Wait.Add(1)
	app.MailChan <- msg
}
