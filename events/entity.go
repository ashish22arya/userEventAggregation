package events

type Events struct {
	UserId    int64  `json:"userId"`
	EventType string `json:"eventType"`
	TimeStamp int64  `json:"timestamp"`
}
