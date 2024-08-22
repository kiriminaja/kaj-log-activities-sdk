package kaj_log_activities_sdk

import (
	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/event"
	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/packages"
	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/requester"
)

type Client struct {
	cfg      Config
	request  requester.Contract
	Events   event.EventContract
	Packages packages.PackageContract
}

func NewClient(cfg Config) *Client {
	request := requester.New()
	return &Client{
		cfg:     cfg,
		request: request,
		Events: event.NewEvent(event.Config{
			URL: cfg.URL,
		}, request),
		Packages: packages.NewPackage(packages.Config{
			URL: cfg.URL,
		}, request),
	}
}
