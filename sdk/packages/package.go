package packages

import (
	"encoding/json"

	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/requester"
)

type packageCase struct {
	request  requester.Contract
	endpoint string
}

func NewPackage(cfg Config, request requester.Contract) PackageContract {
	return &packageCase{
		request:  request,
		endpoint: cfg.URL + "/ex/v1/package",
	}
}

func (p *packageCase) Search(search string) (*ResponseSearchPackage, error) {
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
	response := &ResponseSearchPackage{}
	err = json.Unmarshal(resulsData, &response)
	return response, err
}
