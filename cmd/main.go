package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"

	"Concurrency1/internal/checker"
	"Concurrency1/internal/config"
	"Concurrency1/internal/logger"
	"Concurrency1/pkg/notifier"

	"github.com/joho/godotenv"
)

func main() {
	logger.InitLogger()

	configFile := flag.String("config", "config.json", "path to configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to load config: %v", err))
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	notifyCh := make(chan string)
	doneCh := make(chan bool)

	notifiers := []notifier.Notifier{}

	console := notifier.NewConsoleNotifier()
	notifiers = append(notifiers, console)

	fileNotifier, err := notifier.NewFileNotifier("alerts.log")
	if err == nil {
		notifiers = append(notifiers, fileNotifier)
	}

	telegram := notifier.NewTelegramNotifier("7902861866:AAGERrEP3rL4FyPcUyFDn8pC_TFTB-w5ZKQ", "880921565")
	notifiers = append(notifiers, telegram)

	errw := godotenv.Load()
	if errw != nil {
		fmt.Println("Error loading .env file")
	}

	emailUsername := os.Getenv("EMAIL_USERNAME")
	emailUsername2 := os.Getenv("EMAIL_USERNAME2")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpPort := os.Getenv("SMTP_PORT")

	email := notifier.NewEmailNotifier(emailUsername, emailUsername2, emailPassword, smtpServer, smtpPort)
	notifiers = append(notifiers, email)

	go func() {
		for msg := range notifyCh {
			for _, s := range notifiers {
				s.SendNotification(msg)
			}
		}
		doneCh <- true
	}()

	var wg sync.WaitGroup
	for _, url := range cfg.URLs {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			res, err := checker.CheckURL(ctx, u)
			if err != nil || res.StatusCode >= 400 {
				notifyCh <- fmt.Sprintf("Site %s is down or responding with code %d", u, res.StatusCode)
			} else {
				msg := fmt.Sprintf("Site %s responded successfully in %s", u, res.Duration)
				logger.Info(msg)
				notifyCh <- msg
			}
		}(url)
	}

	wg.Wait()
	close(notifyCh)
	<-doneCh
}
