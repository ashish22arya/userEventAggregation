package main

import (
	"aggregate_events/aggregator"
	"flag"
	"fmt"
)

func main() {
	// Setting up flags
	inputFile := flag.String("i", "", "pass the input json file path")
	outputFile := flag.String("o", "", "pass the output file name")
	updateFlag := flag.Bool("update", false, "pass the update argument to add contribution to last event")

	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Input json file is required !!!")
		return
	}

	if *outputFile == "" {
		fmt.Println("Output file is required !!!")
	}

	if *updateFlag {
		error := aggregator.UpdateSummary(*inputFile, *outputFile)
		if error != nil {
			fmt.Println("Error in updating summary. Error:", error.Error())
			return
		}
	}

	error := aggregator.GenerateSummary(*inputFile, *outputFile)
	if error != nil {
		fmt.Println("Error in generating summary. Error:", error.Error())
		return
	}

	fmt.Println("Summary is written back to output.json file.")
}
