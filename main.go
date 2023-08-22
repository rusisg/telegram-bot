package main

import (
	"flag"
	"log"
	"telegram-bot/telegram"
)

func main() {
	t := mustToken()

	tgClient := telegram.New(token)

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String("token-bot-token", "", "token for access to telegram bot")

	flag.Parse()

	if *token == nil {
		log.Fatal("token is not specified")
	}
	return *token
}
