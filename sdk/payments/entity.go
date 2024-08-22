package payments

type Config struct {
	URL string `json:"url"`
}

//go:generate easytags $GOFILE json

type HistoryData struct {
	EventTrackingID string                   `json:"event_tracking_id"`
	EventName       string                   `json:"event_name"`
	Data            []map[string]interface{} `json:"data"`
}
type ResponseSearch struct {
	Source    string        `json:"source"`
	PaymentID int           `json:"payment_id"`
	TrxID     string        `json:"trx_id"`
	History   []HistoryData `json:"history"`
}
