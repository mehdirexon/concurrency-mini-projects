package models

type Message struct {
	From          string
	FromName      string
	To            string
	Subject       string
	Attachments   []string
	AttachmentMap map[string]string
	Data          any
	DataMap       map[string]any
	Template      string
}
