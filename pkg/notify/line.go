package notify

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const TYPE = "text"

type LineClient struct {
	channelToken string
	client       *http.Client
}

func New(channelToken string) *LineClient {
	return &LineClient{channelToken: channelToken, client: http.DefaultClient}
}

func (l *LineClient) Notify(msg string) error {
	url := "https://api.line.me/v2/bot/message/broadcast"
	ms := Messages{[]Message{{Type: TYPE, Text: msg}}}
	buff, err := json.Marshal(&ms)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(buff))
	req.Header.Set("Authorization", "Bearer "+l.channelToken)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	resp, err := l.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	return nil
}

type Messages struct {
	Messages []Message `json:"messages"`
}

type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
