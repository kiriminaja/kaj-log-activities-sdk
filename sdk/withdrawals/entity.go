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
	Source    string        `json:"source"`
	PackageID int           `json:"package_id"`
	OrderID   string        `json:"order_id"`
	History   []HistoryData `json:"history"`
}
