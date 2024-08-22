package event

import (
	"encoding/json"

	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/requester"
)

type eventCase struct {
	request  requester.Contract
	endpoint string
}

func NewEvent(cfg Config, reqs requester.Contract) EventContract {
	return &eventCase{
		request:  reqs,
		endpoint: cfg.URL + "/ex/v1/event",
	}
}

func (e *eventCase) Upsert(event string, data map[string]interface{}) (map[string]interface{}, error) {
	payload := map[string]interface{}{
		"event": event,
		"data":  data,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	response, err := e.request.POST(e.endpoint, map[string]string{}, body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(response, &result)
	return result, err
}

func (e *eventCase) List() (map[string]interface{}, error) {
	var result Response
	data, err := e.request.GET(e.endpoint, nil)
	if err != nil {
		return nil, err
	}
	err = result.Cast(data, &result)
	if err != nil {
		return result.Data, err
	}
	return result.Data, nil
}
