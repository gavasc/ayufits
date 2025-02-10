package main

import (
	"ayufits/commands"
	"ayufits/data"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tele "gopkg.in/telebot.v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	conf := tele.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(conf)
	if err != nil {
		log.Fatal(err)
		return
	}

	data.CreateDB()

	// Handlers

	b.Handle("/start", func(c tele.Context) error {
		msg := commands.HandleStart(c)

		return c.Send(msg)
	})

	b.Handle("/comi", func(c tele.Context) error {
		resp := commands.HandleComi(c)

		return c.Send(resp)
	})

	b.Start()
}
