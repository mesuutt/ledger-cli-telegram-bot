package bot

import (
	"log"
	"time"

	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)


func Start(token string) {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	h := Handler{Bot: b}
	b.Handle("/alias", h.Alias)
	b.Handle("/start", h.Start)
	b.Handle("/help", h.Help)
	b.Handle(tb.OnText, h.Text)


	logrus.Info("Telegram bot starting")
	b.Start()
}
