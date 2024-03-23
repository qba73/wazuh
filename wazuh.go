package wazuh

import "errors"

type Client struct {
	Token string
}

func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("missing api key")
	}
	return &Client{Token: apiKey}, nil
}
