package notifier

import "log"

type Notifier interface {
	SendNotification(msg string)
}

type ConsoleNotifier struct{}

func NewConsoleNotifier() *ConsoleNotifier {
	return &ConsoleNotifier{}
}

func (n *ConsoleNotifier) SendNotification(msg string) {
	log.Printf("[NOTIFICATION]: %s", msg)
}
