package notifier

import(
	"fmt"
)

// реализация Email-уведомлений
type EmailSender struct {}

func NewEmailSender () *EmailSender {
  return &EmailSender{}
}

func (e *EmailSender) Send(customer string) error {
  fmt.Printf("Email уведомление отправлено клиенту %s\n", customer)
  return nil
}
