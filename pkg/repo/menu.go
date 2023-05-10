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

func getDishesFromBy(from, by int, source []models.Dish) []models.Dish {
	dishes := make([]models.Dish, 0, by)

	if from > len(source) {
		return nil
	}

	for i := from - 1; i < from+by-1; i++ {
		dishes = append(dishes, source[i])
		if i+1 == len(source) {
			break
		}
	}

	return dishes
}

func GetBreakfastDishes(from, by int) []models.Dish {
	menu, err := GetMenu()
	if err != nil {
		return nil
	}

	return getDishesFromBy(from, by, menu.Breakfast)
}

func GetLunchDishes(from, by int) []models.Dish {
	menu, err := GetMenu()
	if err != nil {
		return nil
	}

	return getDishesFromBy(from, by, menu.Lunch)
}

func GetAfternoonSnackDishes(from, by int) []models.Dish {
	menu, err := GetMenu()
	if err != nil {
		return nil
	}

	return getDishesFromBy(from, by, menu.AfternoonSnack)
}

func GetDinnerDishes(from, by int) []models.Dish {
	menu, err := GetMenu()
	if err != nil {
		return nil
	}

	return getDishesFromBy(from, by, menu.Dinner)
}
