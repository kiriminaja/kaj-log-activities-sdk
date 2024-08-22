package kaj_log_activities_sdk

import (
	"github.com/kiriminaja/kaj-log-activities-sdk/event"
	"github.com/kiriminaja/kaj-log-activities-sdk/requester"
)

type Client struct {
	cfg     Config
	events  event.EventContract
	request requester.Contract
}

func NewClient(cfg Config) *Client {
	request := requester.New()
	return &Client{
		cfg:     cfg,
		request: request,
		events: event.NewEvent(event.Config{
			URL: cfg.URL,
		}, request),
	}
}
