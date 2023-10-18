package aggregator

import (
	"aggregate_events/events"
	"aggregate_events/summary"
	"aggregate_events/utils"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

func GenerateSummary(inputFile, outputFile string) error {
	inputEvents, err := utils.ReadEventFromFile(inputFile)
	if err != nil {
		return err
	}

	var events []events.Events
	unmarshalErr := json.Unmarshal(inputEvents, &events)
	if unmarshalErr != nil {
		fmt.Println(unmarshalErr)
		return errors.New("error in reading input json file")
	}

	var allSummary = make(map[string]map[int64]summary.Summary)

	totalEvents := len(events)
	for i := 0; i < totalEvents; i++ {
		eventSummary := summary.GetEventSummary(events[i])

		if _, ok := allSummary[eventSummary.Date]; !ok {
			allSummary[eventSummary.Date] = map[int64]summary.Summary{}
		}

		datewiseSummary := allSummary[eventSummary.Date]
		if _, ok := datewiseSummary[eventSummary.UserId]; !ok {
			datewiseSummary[eventSummary.UserId] = eventSummary
		} else {
			datewiseSummary[eventSummary.UserId] = updateSummary(datewiseSummary[eventSummary.UserId], events[i])
		}

	}

	summary := getSummaryJSON(allSummary)

	err = utils.WriteToFile(outputFile, summary)
	if err != nil {
		return err
	}

	return nil
}

func getSummaryJSON(allSummary map[string]map[int64]summary.Summary) []byte {
	var summaryArr []summary.Summary
	dateArr := []string{}

	// Sorting based on date
	for k := range allSummary {
		dateArr = append(dateArr, k)
	}
	sort.Strings(dateArr)

	dateArrLen := len(dateArr)
	for i := 0; i < dateArrLen; i++ {
		dateSummary := allSummary[dateArr[i]]

		usersArr := []int{}
		for k := range dateSummary {
			usersArr = append(usersArr, int(k))
		}
		sort.Ints(usersArr)

		usersArrLen := len(usersArr)

		for j := 0; j < usersArrLen; j++ {
			summaryArr = append(summaryArr, dateSummary[int64(usersArr[j])])
		}
	}

	summary, _ := json.Marshal(summaryArr)

	return summary
}

func updateSummary(userSummary summary.Summary, newEvent events.Events) summary.Summary {
	switch newEvent.EventType {
	case events.EVENT_TYPE_POST:
		userSummary.Post++
	case events.EVENT_TYPE_LIKE_RECEIVED:
		userSummary.LikeReceived++
	case events.EVENT_TYPE_COMMENT:
		userSummary.Comment++
	}

	return userSummary
}
