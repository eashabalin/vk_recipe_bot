package main

import (
	"fmt"
	"log"
	configs "vk_recipe_bot/pkg/config"
	"vk_recipe_bot/vkbotapi"
)

const (
	ApiVersion = "5.131"
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
		fmt.Println(update.Object.Message.Text)
	}
}
