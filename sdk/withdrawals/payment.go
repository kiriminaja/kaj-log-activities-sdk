package withdrawals

import (
	"encoding/json"

	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/requester"
)

type withdrawalCase struct {
	request  requester.Contract
	endpoint string
}

func NewWithdrawal(cfg Config, request requester.Contract) WithdrawalContract {
	return &withdrawalCase{
		request:  request,
		endpoint: cfg.URL + "/ex/v1/withdrawal",
	}
}

func (p *withdrawalCase) Search(search string) (*ResponseSearch, error) {
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
