package notifier

import(
	"fmt"
)

// реализация SMS-уведомлений
type SMSSender struct {}

func NewSMSSender () *SMSSender {
  return &SMSSender{}
}

func (s *SMSSender) Send(customer string) error{
   fmt.Printf("Sms уведомление отправлено клиенту %s\n", customer)
   return nil
}
