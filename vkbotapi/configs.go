package vkbotapi

import (
	"math/rand"
	"strconv"
)

const (
	APIEndpoint = "https://api.vk.com/method/"
	APIVersion  = "5.131"
	Buffer      = 100
)

type Chattable interface {
	method() string
	params() (Params, error)
}

type LongPollConfig struct {
	GroupID string
}

func (LongPollConfig) method() string {
	return "groups.getLongPollServer"
}

func (c LongPollConfig) params() (Params, error) {
	params := make(Params)

	params.AddNonEmpty("group_id", c.GroupID)
	params.AddNonEmpty("v", APIVersion)

	return params, nil
}

type UpdateConfig struct {
	Ts      string
	Timeout int
	Key     string
}

func (c UpdateConfig) params() (Params, error) {
	params := make(Params)

	params.AddNonEmpty("act", "a_check")
	params.AddNonEmpty("ts", c.Ts)
	params.AddNonZero("wait", c.Timeout)
	params.AddNonEmpty("key", c.Key)

	return params, nil
}

type SendMessageConfig struct {
	UserID   int
	Message  string
	Keyboard *Keyboard
}

func (SendMessageConfig) method() string {
	return "messages.send"
}

func (c SendMessageConfig) params() (Params, error) {
	params := make(Params)

	params.AddNonZero("user_id", c.UserID)
	params.AddNonZero("random_id", int(rand.Int31()))
	params.AddNonEmpty("message", c.Message)
	params.AddNonEmpty("v", APIVersion)
	err := params.AddInterface("keyboard", c.Keyboard)

	return params, err
}

type EventAnswerConfig struct {
	EventID string
	UserID  int
}

func (EventAnswerConfig) method() string {
	return "messages.sendMessageEventAnswer"
}

func (c EventAnswerConfig) params() (Params, error) {
	params := make(Params)

	params.AddNonEmpty("event_id", c.EventID)
	params.AddNonZero("user_id", c.UserID)
	params.AddNonZero("peer_id", c.UserID)
	params.AddNonEmpty("v", APIVersion)

	return params, nil
}

type DeleteMessageConfig struct {
	MessageIDs   []int
	DeleteForAll bool
	GroupID      string
}

func (DeleteMessageConfig) method() string {
	return "messages.delete"
}

func (c DeleteMessageConfig) params() (Params, error) {
	params := make(Params)

	messageIDsStr := ""

	for i := 0; i < len(c.MessageIDs); i++ {
		messageIDsStr += strconv.Itoa(c.MessageIDs[i])
		if i < len(c.MessageIDs)-1 {
			messageIDsStr += ","
		}
	}

	params.AddNonEmpty("message_ids", messageIDsStr)
	params.AddNonEmpty("group_id", c.GroupID)
	params.AddBool("delete_for_all", c.DeleteForAll)
	params.AddNonEmpty("v", APIVersion)

	return params, nil
}

type EditMessageConfig struct {
	PeerID                int
	Message               string
	ConversationMessageID int
	Keyboard              *Keyboard
}

func (c *EditMessageConfig) method() string {
	return "messages.edit"
}

func (c *EditMessageConfig) params() (Params, error) {
	params := make(Params)

	params.AddNonZero("peer_id", c.PeerID)
	params.AddNonEmpty("message", c.Message)
	params.AddNonZero("conversation_message_id", c.ConversationMessageID)
	params.AddNonEmpty("v", APIVersion)
	err := params.AddInterface("keyboard", c.Keyboard)

	return params, err
}
