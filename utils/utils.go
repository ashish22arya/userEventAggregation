package utils

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

const (
	DATEFORMAT           = "2006-01-02"
	PROCESSED_COUNT_FILE = "processed.json"
)

type Processed struct {
	Count int `json:"count"`
}

func ReadFromFile(inputFile string) ([]byte, error) {
	var data []byte
	data, fileErr := os.ReadFile(inputFile)
	if fileErr != nil {
		return data, errors.New("error in finding input file")
	}

	return data, nil
}

func WriteToFile(outputFile string, summary []byte) error {
	os.WriteFile(outputFile, summary, os.ModePerm)

	return nil
}

func GetDateFromEpoch(timestamp int64) string {
	datetime := time.Unix(timestamp, 0)
	return datetime.Format(DATEFORMAT)
}

func GetAlreadyProcessedCount() int {
	data, fileErr := os.ReadFile(PROCESSED_COUNT_FILE)
	if fileErr != nil {
		return 0
	}

	var processed Processed
	unmarshalErr := json.Unmarshal(data, &processed)
	if unmarshalErr != nil {
		return 0
	}

	return processed.Count
}

func SetAlreadyProcessedCount(count int) {
	processed, _ := json.Marshal(Processed{Count: count})
	os.WriteFile(PROCESSED_COUNT_FILE, processed, os.ModePerm)
}
