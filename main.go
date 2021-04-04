package main

import (
	"fmt"
	"log"
	"os"

	crypto_service "github.com/amupxm/go-crypto-bot/crypto_service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	botToken := os.Getenv("BOT_TOKEN")
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.InlineQuery == nil {
			continue
		}
		articleArrays := createInlineQueryResultArticles(update.InlineQuery.ID, update.InlineQuery.Query)
		bot.AnswerInlineQuery(tgbotapi.InlineConfig{InlineQueryID: update.InlineQuery.ID, CacheTime: 0, Results: articleArrays})

	}
}

func createInlineQueryResultArticles(ID string, query string) []interface{} {
	msgArray := []interface{}{}
	results, _ := crypto_service.SearchByName(query)
	for _, item := range results {

		msg := tgbotapi.NewInlineQueryResultArticleMarkdown(item.Name, item.Name, fmt.Sprintf("***%-3s (*%s*)***  \n💲Price: %f\n⌛️last hour: %f%%\n📅last day: %f%%\n🌜last Week: %f%%\n", item.Name, item.Symbol, item.Price, item.PercentChange1h, item.PercentChange24h, item.PercentChange1w))
		msgArray = append(msgArray, msg)
	}
	return msgArray
}
