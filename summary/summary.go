package summary

import (
	"aggregate_events/events"
	"aggregate_events/utils"
)

func GetEventSummary(event events.Events) Summary {
	summary := Summary{
		UserId:       event.UserId,
		Date:         utils.GetDateFromEpoch(event.TimeStamp),
		Post:         0,
		LikeReceived: 0,
		Comment:      0,
	}

	switch event.EventType {
	case events.EVENT_TYPE_POST:
		summary.Post++
	case events.EVENT_TYPE_LIKE_RECEIVED:
		summary.LikeReceived++
	case events.EVENT_TYPE_COMMENT:
		summary.Comment++
	}

	return summary
}
