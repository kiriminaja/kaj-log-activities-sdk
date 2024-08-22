package event

import "github.com/kiriminaja/kaj-log-activities-sdk/requester"

type eventCase struct {
	request requester.Contract
	config  Config
}

func NewEvent(cfg Config, reqs requester.Contract) EventContract {
	return &eventCase{
		request: reqs,
		config:  cfg,
	}
}

func (e *eventCase) Upsert() {

}
