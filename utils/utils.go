package utils

import (
	"errors"
	"os"
	"time"
)

const (
	DATEFORMAT = "2006-01-02"
)

func ReadEventFromFile(inputFile string) ([]byte, error) {
	var eventData []byte
	eventData, fileErr := os.ReadFile(inputFile)
	if fileErr != nil {
		return eventData, errors.New("error in finding input file")
	}

	return eventData, nil
}

func WriteToFile(outputFile string, summary []byte) error {
	os.WriteFile(outputFile, summary, os.ModePerm)

	return nil
}

func GetDateFromEpoch(timestamp int64) string {
	datetime := time.Unix(timestamp, 0)
	return datetime.Format(DATEFORMAT)
}
