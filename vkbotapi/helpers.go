package vkbotapi

func NewMessage(userID int, text string) SendMessageConfig {
	return SendMessageConfig{
		UserID:  userID,
		Message: text,
	}
}
