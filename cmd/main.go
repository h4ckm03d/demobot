package main

import (
	"os"
	"time"

	"github.com/h4ckm03d/demobot"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	b := demobot.NewWebhookBot(os.Getenv("TELEGRAM_TOKEN"))
	b.Setup()
	b.Bot.Poller = &tb.LongPoller{Timeout: 10 * time.Second}
	b.Bot.Start()
}
