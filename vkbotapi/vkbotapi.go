package vkbotapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type VKBotAPI struct {
	token   string
	groupID string
	client  *http.Client
	Debug   bool
	Buffer  int

	apiEndpoint      string
	longPollEndpoint string
}

func NewVKBotAPI(token, groupID string, debug bool) *VKBotAPI {
	return &VKBotAPI{
		token:       token,
		groupID:     groupID,
		client:      &http.Client{},
		apiEndpoint: APIEndpoint,
		Debug:       debug,
		Buffer:      Buffer,
	}
}

func (b *VKBotAPI) GetUpdatesChan(config UpdateConfig) (chan Update, error) {
	longPollConfig := LongPollConfig{GroupID: b.groupID}

	longPollResponse, err := b.GetLongPollServer(longPollConfig)
	if err != nil {
		return nil, err
	}

	config.Ts = longPollResponse.Response.Ts
	config.Key = longPollResponse.Response.Key

	ch := make(chan Update, b.Buffer)

	go func() {
		for {
			updateResponse, err := b.GetUpdates(longPollResponse.Response.Server, config)
			if err != nil {
				log.Println(err)
				log.Println("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)

				continue
			}

			for _, update := range updateResponse.Updates {
				ch <- update
			}
			config.Ts = updateResponse.Ts
		}
	}()

	return ch, nil
}

func (b *VKBotAPI) GetLongPollServer(config LongPollConfig) (*LongPollResponse, error) {
	resp, err := b.Request(config)
	if err != nil {
		return nil, err
	}

	var longPollResponse LongPollResponse

	err = json.Unmarshal(resp.Result, &longPollResponse)
	if err != nil {
		return nil, err
	}

	return &longPollResponse, nil
}

func (b *VKBotAPI) GetUpdates(longPollServer string, config UpdateConfig) (*UpdateResponse, error) {
	params, err := config.params()
	if err != nil {
		return nil, err
	}

	resp, err := b.RequestURL(longPollServer, params)
	if err != nil {
		return nil, err
	}
	updateResponse := UpdateResponse{}

	err = json.Unmarshal(resp.Result, &updateResponse)
	if err != nil {
		return nil, err
	}

	return &updateResponse, nil
}

func buildParams(in Params) url.Values {
	if in == nil {
		return url.Values{}
	}

	out := url.Values{}

	for key, value := range in {
		out.Set(key, value)
	}

	return out
}

func (b *VKBotAPI) Request(c Chattable) (*APIResponse, error) {
	params, err := c.params()
	if err != nil {
		return nil, err
	}

	u := b.apiEndpoint + c.method()

	return b.RequestURL(u, params)
}

func (b *VKBotAPI) RequestURL(url string, params Params) (*APIResponse, error) {
	if b.Debug {
		log.Printf("Endpoint: %s, params: %v\n", url, params)
	}

	values := buildParams(params)

	method := url + "?" + values.Encode()

	req, err := http.NewRequest(http.MethodPost, method, nil)

	req.Header.Set("Authorization", "Bearer "+b.token)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := b.client.Do(req)
	if err != nil {
		return nil, err
	}

	rawData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	errorResponse := ErrorResponse{}

	err = json.Unmarshal(rawData, &errorResponse)
	if err != nil {
		return nil, err
	}

	errorMessage := errorResponse.Error.ErrorMessage
	errorCode := errorResponse.Error.ErrorCode

	if errorMessage != "" {
		return nil, errors.New(fmt.Sprintf("Error code %d. %s\n", errorCode, errorMessage))
	}

	return &APIResponse{
		Result:     rawData,
		StatusCode: resp.StatusCode,
	}, nil
}
