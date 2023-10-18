# userEventAggregation
User event aggregation CLI for SugarBox

Task: 

Part 1: Write a command-line utility that reads a list of user events from a JSON file, aggregates the events, and writes the daily summary reports back to another JSON file.

Steps To Exectute:
Step 1: Build the go Program
$ go build

Step 2: Put the input.json file at the project root.(Already present for testing)

Step 3: Run the program with command line arguments
$ ./aggregate_events -i input.json -o output.json


Part 2: Extend your utility to support real-time aggregation. Each time a new event is added to the input JSON file, the utility should update the corresponding daily summary report without reprocessing all the previous events.

Here, Assuming that the command line with --update flag will be executed every time a new event is added to the input.json file.

To execute this, Append --update flag to the step 3 abvove:
$ ./aggregate_events -i input.json -o output.json --update