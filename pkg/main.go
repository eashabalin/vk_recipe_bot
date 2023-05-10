package main

import (
	"fmt"
	"log"
	bot2 "vk_recipe_bot/pkg/bot"
	configs "vk_recipe_bot/pkg/config"
	"vk_recipe_bot/vkbotapi"
)

func main() {
	cfg, err := configs.GetConfig()
	if err != nil {
		panic(err)
	}

	bot := vkbotapi.NewVKBotAPI(cfg.Token, cfg.GroupID, true)

	config := vkbotapi.UpdateConfig{
		Timeout: 25,
	}

	updates, err := bot.GetUpdatesChan(config)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {

		fmt.Println(update)

		if update.IsMessageNew() {
			err = bot2.HandleMessage(bot, update.MessageNew())
			if err != nil {
				fmt.Println(err)
			}
		}
		if update.IsMessageEvent() {
			err = bot2.HandleEvent(bot, update.MessageEvent())
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
