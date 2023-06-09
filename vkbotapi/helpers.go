package vkbotapi

func NewMessage(userID int, text string) SendMessageConfig {
	return SendMessageConfig{
		UserID:  userID,
		Message: text,
	}
}

func newKeyboard(inline bool, buttonRows ...[]Button) *Keyboard {
	return &Keyboard{
		Inline:  inline,
		Buttons: buttonRows,
	}
}

func NewKeyboardButton(text string) Button {
	return Button{
		Action: Action{
			Type:  Text,
			Label: text,
		},
	}
}

func NewKeyboardCallbackButton(text string, command string) Button {
	return Button{
		Action: Action{
			Type:    Callback,
			Label:   text,
			Payload: Payload{Command: command},
		},
	}
}

func NewKeyboardButtonRow(buttons ...Button) []Button {
	return buttons
}

func NewKeyboard(buttonRows ...[]Button) *Keyboard {
	return newKeyboard(false, buttonRows...)
}

func NewInlineKeyboard(buttonRows ...[]Button) *Keyboard {
	return newKeyboard(true, buttonRows...)
}

func NewDeleteMessage(deleteForAll bool, groupID string, messageIDs ...int) *DeleteMessageConfig {
	return &DeleteMessageConfig{
		MessageIDs:   messageIDs,
		DeleteForAll: deleteForAll,
		GroupID:      groupID,
	}
}

func NewEditMessage(userID int, conversationMessageID int, message string) *EditMessageConfig {
	return &EditMessageConfig{
		PeerID:                userID,
		Message:               message,
		ConversationMessageID: conversationMessageID,
	}
}

func NewEventAnswer(eventID string, userID int) *EventAnswerConfig {
	return &EventAnswerConfig{
		EventID: eventID,
		UserID:  userID,
	}
}
