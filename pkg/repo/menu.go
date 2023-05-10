package repo

import (
	"encoding/json"
	"os"
	"strings"
	"vk_recipe_bot/pkg/models"
)

func GetMenu() (*models.Menu, error) {
	data, err := os.ReadFile("pkg/repo/menu.json")
	if err != nil {
		return nil, err
	}

	var menu models.Menu

	err = json.Unmarshal(data, &menu)
	if err != nil {
		return nil, err
	}
	return &menu, nil
}

func FindDish(name string) *models.Dish {
	menu, err := GetMenu()
	if err != nil {
		return nil
	}

	for _, d := range menu.Breakfast {
		if strings.ToLower(d.Name) == strings.ToLower(name) {
			return &d
		}
	}
	for _, d := range menu.Lunch {
		if strings.ToLower(d.Name) == strings.ToLower(name) {
			return &d
		}
	}
	for _, d := range menu.AfternoonSnack {
		if strings.ToLower(d.Name) == strings.ToLower(name) {
			return &d
		}
	}
	for _, d := range menu.Dinner {
		if strings.ToLower(d.Name) == strings.ToLower(name) {
			return &d
		}
	}
	return nil
}
