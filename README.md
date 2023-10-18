# userEventAggregation
User event aggregation CLI for SugarBox

Task: 
Part 1: Write a command-line utility that reads a list of user events from a JSON file, aggregates the events, and writes the daily summary reports back to another JSON file.

Steps To Exectute:
Step 1: Build the go Program
$ go build

Step 2: Put the input.json file at the project root.(Already present for testing)

Step 2: Run the program with command line arguments
$ ./aggregate_events -i input.json -o output.json