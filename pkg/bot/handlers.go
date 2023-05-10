package bot

import (
	"fmt"
	"vk_recipe_bot/pkg/repo"
	"vk_recipe_bot/vkbotapi"
)

func HandleMessage(b *vkbotapi.VKBotAPI, m *vkbotapi.Message) error {
	menu, err := repo.GetMenu()
	if err != nil {
		return err
	}

	if m.Text == "Завтрак" {
		msg := vkbotapi.NewMessage(m.FromID, "Круто! Какое блюдо ты хотел бы приготовить?😋")

		buttonRows := make([][]vkbotapi.Button, 6)

		for i := 0; i < 5; i++ {
			buttonRows[i] = append(buttonRows[i], vkbotapi.NewKeyboardCallbackButton(menu.Breakfast[i].Name, menu.Breakfast[i].Name))
		}
		buttonRows[5] = append(buttonRows[5], vkbotapi.NewKeyboardButton("Ещё"))

		keyboard := vkbotapi.NewInlineKeyboard(buttonRows...)

		msg.Keyboard = keyboard

		err = b.Send(msg)
		if err != nil {
			return err
		}
	}
	if m.Text == "Начать" {
		msg := vkbotapi.NewMessage(m.FromID, "Здорово! Какой приём пищи тебя интересует: завтрак, обед или ужин?🤔")

		keyboard := vkbotapi.NewInlineKeyboard(
			vkbotapi.NewKeyboardButtonRow(
				vkbotapi.NewKeyboardButton("Завтрак")),
			vkbotapi.NewKeyboardButtonRow(
				vkbotapi.NewKeyboardButton("Обед")),
			vkbotapi.NewKeyboardButtonRow(
				vkbotapi.NewKeyboardButton("Полдник")),
			vkbotapi.NewKeyboardButtonRow(
				vkbotapi.NewKeyboardButton("Ужин")),
		)

		msg.Keyboard = keyboard

		err := b.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}

func HandleEvent(b *vkbotapi.VKBotAPI, m *vkbotapi.MessageEventObject) error {
	dish := repo.FindDish(m.Payload.Command)
	if dish != nil {
		msg := vkbotapi.NewMessage(m.UserID, dish.Ingredients+"\n"+dish.Recipe)
		err := b.Send(msg)
		if err != nil {
			return err
		}
	}
	answer := vkbotapi.EventAnswerConfig{EventID: m.EventID, UserID: m.UserID}
	err := b.Send(answer)
	if err != nil {
		return err
	}

	edit := vkbotapi.NewEditMessage(m.UserID, m.ConversationMessageID, "лалалал")

	keyboard := vkbotapi.NewInlineKeyboard(
		vkbotapi.NewKeyboardButtonRow(
			vkbotapi.NewKeyboardButton("Завтрак")),
		vkbotapi.NewKeyboardButtonRow(
			vkbotapi.NewKeyboardButton("Обед")),
		vkbotapi.NewKeyboardButtonRow(
			vkbotapi.NewKeyboardButton("Полдник")),
		vkbotapi.NewKeyboardButtonRow(
			vkbotapi.NewKeyboardButton("Ужин")),
	)

	edit.Keyboard = keyboard

	err = b.Send(edit)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
