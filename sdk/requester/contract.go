package requester

import (
	"io"
	"net/http"

	"github.com/kiriminaja/kaj-log-activities-sdk/sdk/presentation"
)

type Contract interface {
	GET(urlBase string, header, params map[string]string) (presentation.Response, error)
	DELETE(url string, header map[string]string) (presentation.Response, error)
	POST(url string, header map[string]string, payload []byte) (presentation.Response, error)
	PUT(url string, header map[string]string, payload []byte) (presentation.Response, error)
	RAW(method, url string, body io.Reader) (*http.Request, error)
	WithBasicPOST(url string, header map[string]string, payload []byte, username, password string) (presentation.Response, error)
}
