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
	GroupID int         `json:"group_id"`
	Type    UpdateType  `json:"type"`
	EventID string      `json:"event_id"`
	Object  interface{} `json:"object"`
}

func (u Update) IsMessageNew() bool {
	if u.Type == MessageNew {
		return true
	}
	return false
}

func (u Update) IsMessageEvent() bool {
	if u.Type == MessageEvent {
		return true
	}
	return false
}

func (u Update) MessageNew() *Message {
	m, ok := u.Object.(MessageNewObject)
	if ok {
		return &m.Message
	} else {
		return nil
	}
}

func (u Update) MessageEvent() *MessageEventObject {
	m, ok := u.Object.(MessageEventObject)
	if ok {
		return &m
	} else {
		return nil
	}
}

type UpdateType string

const (
	MessageNew   UpdateType = "message_new"
	MessageReply UpdateType = "message_reply"
	MessageEvent UpdateType = "message_event"
)

type MessageNewObject struct {
	Message Message `json:"message"`
}

type Message struct {
	FromID   int       `json:"from_id"`
	ID       int       `json:"id"`
	Text     string    `json:"text"`
	Keyboard *Keyboard `json:"keyboard"`
}

type Error struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_msg"`
}

type ErrorResponse struct {
	Error Error `json:"error"`
}

type ActionType string

const (
	Text     ActionType = "text"
	OpenLink ActionType = "open_link"
	Callback ActionType = "callback"
)

type Action struct {
	Type    ActionType `json:"type"`
	Link    string     `json:"link,omitempty"`
	Label   string     `json:"label"`
	Payload Payload    `json:"payload,omitempty"`
}

type Button struct {
	Action Action `json:"action"`
}

type Keyboard struct {
	Inline  bool       `json:"inline"`
	Buttons [][]Button `json:"buttons"`
}

type MessageEventObject struct {
	ConversationMessageID int     `json:"conversation_message_id"`
	UserID                int     `json:"user_id"`
	PeerID                int     `json:"peer_id"`
	EventID               string  `json:"event_id"`
	Payload               Payload `json:"payload"`
}

type Payload struct {
	Command string `json:"command"`
}
