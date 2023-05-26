package main

import (
	"twint-zero/Core"
	"twint-zero/InputParser"
	"time"
	"fmt"
	"strconv"
)

func main() {
	Arguments := InputParser.InputParser()
	increment := 4
	var sinceDateList []string
	var untilDateList []string
	startDate,err := strconv.Atoi(Arguments.StartDate)
	if(err != nil){
		fmt.Println(err)
	}
	start := time.Date(startDate, 12, 31, 0, 0, 0, 0, time.UTC)
	start = start
	end := time.Now()
	for d := start; d.After(end.AddDate(0, 0, -increment)) == false; d = d.AddDate(0, 0, increment){
		sinceDateList = append(sinceDateList, d.Format("2006-01-02"))
		fmt.Println(d.Format("2006-01-02"))
	}
	untilStart := start.AddDate(0, 0, increment - 1)
	untilEnd := time.Now()
	for d := untilStart; d.After(untilEnd) == false; d = d.AddDate(0, 0, increment) {
		if(d.AddDate(0, 0, increment).After(untilEnd)){
			d = untilEnd
		}
		untilDateList = append(untilDateList, d.Format("2006-01-02"))
		fmt.Println(d.Format("2006-01-02"))
	}
	for i := 0; i < len(sinceDateList); i = i + 1{
		query := Arguments.Query + sinceDateList[i] + " until:" + untilDateList[i]	
    	Core.Main(&(query), &(Arguments.Instance), &(Arguments.Format), &(Arguments.Name))
    }
}	
