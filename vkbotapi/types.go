package vkbotapi

import (
	"encoding/json"
)

type APIResponse struct {
	Result     json.RawMessage
	StatusCode int
}

type LongPollResponse struct {
	Response struct {
		Key    string `json:"key"`
		Server string `json:"server"`
		Ts     string `json:"ts"`
	} `json:"response"`
}

type UpdateResponse struct {
	Ts      string   `json:"ts"`
	Updates []Update `json:"updates"`
}

type Update struct {
	GroupID int        `json:"group_id"`
	Type    UpdateType `json:"type"`
	EventID string     `json:"event_id"`
	Object  `json:"object"`
}

type UpdateType string

const (
	MessageNew   UpdateType = "message_new"
	MessageReply UpdateType = "message_reply"
)

type Object struct {
	Message Message `json:"message"`
}

type Message struct {
	FromID   int      `json:"from_id"`
	ID       int      `json:"id"`
	Text     string   `json:"text"`
	Keyboard Keyboard `json:"keyboard"`
}

type Error struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_msg"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}
