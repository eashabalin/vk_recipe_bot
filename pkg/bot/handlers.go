package bot

import (
	"fmt"
	"strconv"
	"strings"
	"vk_recipe_bot/pkg/models"
	"vk_recipe_bot/pkg/repo"
	"vk_recipe_bot/vkbotapi"
)

const (
	mealsText  = "Здорово! Какой приём пищи тебя интересует: завтрак, обед или ужин?🤔"
	dishesText = "Какое блюдо выберешь?😋"
	startText  = "Чтобы начать, отправь сообщение \"Начать\""
)

func HandleMessage(b *vkbotapi.VKBotAPI, m *vkbotapi.Message) error {

	if m.Text == "Начать" {
		msg := vkbotapi.NewMessage(m.FromID, mealsText)

		keyboard := StartKeyboard()

		msg.Keyboard = keyboard

		err := b.Send(msg)
		if err != nil {
			return err
		}
		return nil
	} else {
		msg := vkbotapi.NewMessage(m.FromID, startText)

		err := b.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func StartKeyboard() *vkbotapi.Keyboard {
	return vkbotapi.NewInlineKeyboard(
		vkbotapi.NewKeyboardButtonRow(
			vkbotapi.NewKeyboardCallbackButton("Завтрак", "breakfast 1 3")),
		vkbotapi.NewKeyboardButtonRow(
			vkbotapi.NewKeyboardCallbackButton("Обед", "lunch 1 3")),
		vkbotapi.NewKeyboardButtonRow(
			vkbotapi.NewKeyboardCallbackButton("Полдник", "afternoon_snack 1 3")),
		vkbotapi.NewKeyboardButtonRow(
			vkbotapi.NewKeyboardCallbackButton("Ужин", "dinner 1 3")),
	)
}

type MealOption struct {
	Meal string
	From int
	By   int
}

func NewMealOption(in string) (new MealOption, ok bool) {
	f := strings.Fields(in)
	if len(f) != 3 {
		return MealOption{}, false
	}
	if !(f[0] == "breakfast" || f[0] == "lunch" || f[0] == "afternoon_snack" || f[0] == "dinner") {
		return MealOption{}, false
	}
	from, err := strconv.Atoi(f[1])
	if err != nil {
		return MealOption{}, false
	}
	by, err := strconv.Atoi(f[2])
	if err != nil {
		return MealOption{}, false
	}
	return MealOption{
		Meal: f[0],
		From: from,
		By:   by,
	}, true
}

func HandleEvent(b *vkbotapi.VKBotAPI, m *vkbotapi.MessageEventObject) error {
	answer := vkbotapi.NewEventAnswer(m.EventID, m.UserID)
	err := b.Send(answer)
	if err != nil {
		return err
	}

	dish := repo.FindDish(m.Payload.Command)

	if dish != nil {
		msg := vkbotapi.NewMessage(
			m.UserID,
			fmt.Sprintf("%s\n%s\n%s\n", dish.Name, dish.Ingredients, dish.Recipe),
		)

		keyboard := vkbotapi.NewInlineKeyboard(
			vkbotapi.NewKeyboardButtonRow(
				vkbotapi.NewKeyboardCallbackButton("Поискать ещё", "start")))

		msg.Keyboard = keyboard

		err = b.Send(msg)
		if err != nil {
			return err
		}
	}

	if m.Payload.Command == "from_start" {
		editMsg := vkbotapi.NewEditMessage(m.UserID, m.ConversationMessageID, mealsText)
		keyboard := StartKeyboard()
		editMsg.Keyboard = keyboard
		err = b.Send(editMsg)
		if err != nil {
			return err
		}
	}
	if m.Payload.Command == "start" {
		msg := vkbotapi.NewMessage(m.UserID, mealsText)
		keyboard := StartKeyboard()
		msg.Keyboard = keyboard
		err = b.Send(msg)
		if err != nil {
			return err
		}
	}

	mealOption, ok := NewMealOption(m.Payload.Command)
	if ok {
		menu, err := repo.GetMenu()
		if err != nil {
			return err
		}

		return ShowDishes(b, &mealOption, m, menu)
	}

	return nil
}

func ShowDishes(b *vkbotapi.VKBotAPI, mealOption *MealOption, m *vkbotapi.MessageEventObject, menu *models.Menu) error {
	dishes := repo.GetBreakfastDishes(mealOption.From, mealOption.By)
	maxLength := len(menu.Breakfast)

	if mealOption.Meal == "lunch" {
		dishes = repo.GetLunchDishes(mealOption.From, mealOption.By)
		maxLength = len(menu.Lunch)
	}
	if mealOption.Meal == "afternoon_snack" {
		dishes = repo.GetAfternoonSnackDishes(mealOption.From, mealOption.By)
		maxLength = len(menu.AfternoonSnack)
	}
	if mealOption.Meal == "dinner" {
		dishes = repo.GetDinnerDishes(mealOption.From, mealOption.By)
		maxLength = len(menu.Dinner)
	}

	l := len(dishes)
	more := true
	if mealOption.From+mealOption.By-1 > maxLength {
		more = false
	}
	btnRows := make([][]vkbotapi.Button, l+1, l+2)

	for i := 0; i < l; i++ {
		btnRows[i] = append(btnRows[i], vkbotapi.NewKeyboardCallbackButton(dishes[i].Name, dishes[i].Name))
	}

	if maxLength > mealOption.By {
		btnRows = append(btnRows, []vkbotapi.Button{})
		if more {
			btnRows[l] = append(
				btnRows[l],
				vkbotapi.NewKeyboardCallbackButton(
					"Ещё",
					fmt.Sprintf("%s %d %d", mealOption.Meal, mealOption.From+mealOption.By, mealOption.By),
				),
			)
		} else {
			btnRows[l] = append(
				btnRows[l],
				vkbotapi.NewKeyboardCallbackButton(
					"Сначала",
					fmt.Sprintf("%s %d %d", mealOption.Meal, 1, mealOption.By),
				),
			)
		}
	}
	btnRows[len(btnRows)-1] = append(btnRows[len(btnRows)-1], vkbotapi.NewKeyboardCallbackButton(
		"Назад",
		"from_start",
	))

	keyboard := vkbotapi.NewInlineKeyboard(btnRows...)
	editMsg := vkbotapi.NewEditMessage(m.UserID, m.ConversationMessageID, dishesText)
	editMsg.Keyboard = keyboard
	err := b.Send(editMsg)
	if err != nil {
		return err
	}
	return nil
}
