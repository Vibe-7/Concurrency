package notifier

import (
	"log"
	"os"
)

type FileNotifier struct {
	logger *log.Logger
}

func NewFileNotifier(filePath string) (*FileNotifier, error) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileNotifier{
		logger: log.New(f, "", log.LstdFlags|log.Lmicroseconds),
	}, nil
}

func (n *FileNotifier) SendNotification(msg string) {
	n.logger.Printf("[NOTIFICATION]: %s", msg)
}
