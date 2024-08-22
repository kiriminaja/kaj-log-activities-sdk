package event

type EventContract interface {
	Upsert(event string, data map[string]interface{}) (map[string]interface{}, error)
	List() (map[string]interface{}, error)
}
