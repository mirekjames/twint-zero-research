package Core

import (
	"net/url"
	// "fmt"
)

var (
	condition bool   = true
	cursor    string = ""
)

func Main(Query *string, Instance *string, Format *string, Name *string) {
	(*Query) = url.QueryEscape(*Query)
	cursor = ""
	condition = true
	for condition {
		condition = Scrape(Request(Query, Instance, &cursor), Name, Format, &cursor)
	}
}
