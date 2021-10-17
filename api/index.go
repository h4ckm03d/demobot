package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	b := NewWebhookBot(os.Getenv("TELEGRAM_TOKEN"))
	b.Setup()

	var u tb.Update

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	if err = json.Unmarshal(body, &u); err == nil {
		b.Bot.ProcessUpdate(u)
	}
}
