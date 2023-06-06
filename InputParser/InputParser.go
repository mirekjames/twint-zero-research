package InputParser

import (
	"flag"
	"os"
)

type Arguments struct {
	Query    string
	Instance string
	Format   string
	Name     string
	StartDate string
	EndDate string
	// DateIncrement int
}

var arguments *Arguments = new(Arguments)

func InputParser() *Arguments {

	flag.StringVar(&(arguments.Query), "Query", "", "Specify search query.")
	flag.StringVar(&(arguments.Instance), "Instance", "nitter.nl", "Specify instance to get data from.")
	flag.StringVar(&(arguments.Format), "Format", "csv", "Specify the return format: csv (default), or json.")
	flag.StringVar(&(arguments.Name), "Name", "tweets", "Specify the filename (without file extension).")
	flag.StringVar(&(arguments.StartDate), "StartDate", "", "Specify the start date in format YYYY-MM-DD")
	flag.StringVar(&(arguments.EndDate), "EndDate", "", "Specify the end date in format YYYY-MM-DD")
	// flag.IntVar(&(arguments.DateIncrement), "DateIncrement", "1", "Specify how many days in include in each scrape request.")
	flag.Parse()

	if (*arguments).Query == "" || !ValidateFormatArgument(arguments) {
		flag.Usage()
		os.Exit(1)
	}

	return arguments
}

func ValidateFormatArgument(arguments *Arguments) bool {
	format := (*arguments).Format
	return format == "" || format == "csv" || format == "json"
}
