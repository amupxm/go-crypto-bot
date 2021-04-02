package main

import (
	"fmt"
	"log"

	crypto_service "github.com/amupxm/go-crypto-bot/crypto_service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1415691028:AAF0oHXfu-Xdfl6Q00Fqi9JULd6gh-hAYoQ")
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
		log.Println(update.InlineQuery.Query)
		articleArrays := createInlineQueryResultArticles(update.InlineQuery.ID, update.InlineQuery.Query)
		bot.AnswerInlineQuery(tgbotapi.InlineConfig{InlineQueryID: update.InlineQuery.ID, CacheTime: 0, Results: articleArrays})
	}
}

func createInlineQueryResultArticles(ID string, query string) []interface{} {
	msgArray := []interface{}{}
	results, _ := crypto_service.SearchByName(query)
	for _, item := range results {
		msg := tgbotapi.NewInlineQueryResultArticle(item.Name, item.Name, fmt.Sprintf("%-3s (%s)  \n💲Price: %f\n⌛️last hour: %f%%\n📅last day: %f%%\n🌜last Week: %f%%\n", item.Name, item.Symbol, item.Price, item.PercentChange1h, item.PercentChange24h, item.PercentChange1w))
		msgArray = append(msgArray, msg)
	}
	return msgArray
}
