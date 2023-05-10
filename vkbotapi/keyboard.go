package vkbotapi

type ActionType string

const (
	Text     ActionType = "text"
	OpenLink ActionType = "open_link"
)

type Action struct {
	Type  ActionType `json:"type"`
	Link  string     `json:"link,omitempty"`
	Label string     `json:"label"`
}

type Button struct {
	Action Action `json:"action"`
}

type Keyboard struct {
	Inline  bool       `json:"inline"`
	Buttons [][]Button `json:"buttons"`
}
