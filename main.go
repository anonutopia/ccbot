package main

import (
	"gopkg.in/tucnak/telebot.v2"
	tb "gopkg.in/tucnak/telebot.v2"
)

var b *telebot.Bot

func main() {
	b = initTelegramBot()

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, string(m.Text))
	})

	b.Start()
}
