package packages

import "github.com/kiriminaja/kaj-log-activities-sdk/sdk/requester"

type packageCase struct {
	request  requester.Contract
	endpoint string
}

func NewPackage(cfg Config, request requester.Contract) PackageContract {
	return &packageCase{
		request:  request,
		endpoint: cfg.URL + "/ex/v1/package/",
	}
}
