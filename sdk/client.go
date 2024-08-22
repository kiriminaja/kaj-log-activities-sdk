package kaj_log_activities_sdk

import (
	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/events"
	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/packages"
	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/payments"
	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/requester"
	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/withdrawals"
)

type Client struct {
	cfg         Config
	request     requester.Contract
	Events      events.EventContract
	Packages    packages.PackageContract
	Payments    payments.PaymentsContract
	Withdrawals withdrawals.WithdrawalContract
}

func NewClient(cfg Config) *Client {
	request := requester.New()
	return &Client{
		cfg:     cfg,
		request: request,
		Events: events.NewEvent(events.Config{
			URL: cfg.URL,
		}, request),
		Packages: packages.NewPackage(packages.Config{
			URL: cfg.URL,
		}, request),
		Payments: payments.NewPayment(payments.Config{
			URL: cfg.URL,
		}, request),

		Withdrawals: withdrawals.NewWithdrawal(withdrawals.Config{
			URL: cfg.URL,
		}, request),
	}
}
