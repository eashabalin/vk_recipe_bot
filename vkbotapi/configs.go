package vkbotapi

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
