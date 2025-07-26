package notifier

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type TelegramNotifier struct {
	BotToken string
	ChatID   string
}

func NewTelegramNotifier(botToken, chatID string) *TelegramNotifier {
	return &TelegramNotifier{
		BotToken: botToken,
		ChatID:   chatID,
	}
}

func (t *TelegramNotifier) SendNotification(msg string) {
	url := fmt.Sprintf(
		"https://api.telegram.org/bot%s/sendMessage?chat_id=%s&text=%s",
		t.BotToken, t.ChatID, url.QueryEscape(msg),
	)

	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Failed to send Telegram message: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Telegram API error: %s", string(body))
	}
}
