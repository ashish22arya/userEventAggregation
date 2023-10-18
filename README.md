# userEventAggregation
User event aggregation CLI for SugarBox

Task: 

Part 1: Write a command-line utility that reads a list of user events from a JSON file, aggregates the events, and writes the daily summary reports back to another JSON file.

Steps To Exectute:
Step 1: Build the go Program (Pre-requisite: Required go 1.21.1)
$ go build

Step 2: Put the input.json file at the project root.(Already present for testing)

Step 3: Run the program with command line arguments
$ ./aggregate_events -i input.json -o output.json


Part 2: Extend your utility to support real-time aggregation. Each time a new event is added to the input JSON file, the utility should update the corresponding daily summary report without reprocessing all the previous events.

To execute this, Append --update flag to the step 3 above:
$ ./aggregate_events -i input.json -o output.json --update


Cases handled:
Case 1: When ever the command line triggered after adding any 'x' number of new events, it will handle correctly considering all new added elements.

Case 2: When the input.json file gets reset or have events less that number of event it aggregated in previous run, it will discard the previous output aggregation and re-compute from the very first event.