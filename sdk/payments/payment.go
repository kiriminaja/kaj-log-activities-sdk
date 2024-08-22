package payments

import (
	"encoding/json"

	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/requester"
)

type paymentsCase struct {
	request  requester.Contract
	endpoint string
}

func NewPayment(cfg Config, request requester.Contract) PaymentsContract {
	return &paymentsCase{
		request:  request,
		endpoint: cfg.URL + "/ex/v1/payment",
	}
}

func (p *paymentsCase) Search(search string) (*ResponseSearch, error) {
	enpoint := p.endpoint + "/search"
	params := map[string]string{
		"search": search,
	}
	result, err := p.request.GET(enpoint, nil, params)
	if err != nil {
		return nil, err
	}
	resulsData, err := json.Marshal(result.Data)
	if err != nil {
		return nil, err
	}
	response := &ResponseSearch{}
	err = json.Unmarshal(resulsData, &response)
	return response, err
}
