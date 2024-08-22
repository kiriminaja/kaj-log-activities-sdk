package event

import "encoding/json"

type Config struct {
	URL string
}

type Response struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Data    map[string]interface{}
}

func (r *Response) Cast(response []byte, val interface{}) error {
	return json.Unmarshal(response, &r)
}
