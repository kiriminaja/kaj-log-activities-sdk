package withdrawals

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
	Source       string        `json:"source"`
	WithdrawalID int           `json:"withdrawal_id"`
	RefNum       string        `json:"ref_num"`
	History      []HistoryData `json:"history"`
}
