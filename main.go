package main

import (
	"twint-zero/Core"
	"twint-zero/InputParser"
	"time"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	Arguments := InputParser.InputParser()
	start := time.Date(2007, 01, 02, 0, 0, 0, 0, time.UTC)
	end := time.Now()

	if Arguments.StartDate != "" {
		startDateComponents := strings.Split(Arguments.StartDate, "-")
		startYear, err := strconv.Atoi(startDateComponents[0])
		if err != nil{
			fmt.Println(err)
		}
		startMonth, err := strconv.Atoi(startDateComponents[1])
		if err != nil{
			fmt.Println(err)
		}
		startDay, err := strconv.Atoi(startDateComponents[2])
			if err != nil{
			fmt.Println(err)
		}

		start = time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	} else {
		start = time.Date(2007, 01, 02, 0, 0, 0, 0, time.UTC)
	}

	if Arguments.EndDate != "" {
		endDateComponents := strings.Split(Arguments.EndDate, "-")
		endYear, err := strconv.Atoi(endDateComponents[0])
		if err != nil {
			fmt.Println(err)
		}
		endMonth, err := strconv.Atoi(endDateComponents[1])
		if err != nil{
			fmt.Println(err)
		}
		endDay, err := strconv.Atoi(endDateComponents[2])
		if err != nil{
			fmt.Println(err)
		}	

		end = time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)
	} else {
		end = time.Now()
	}

	var sinceDateList []string
	var untilDateList []string
	increment := 4
	for d := start; d.After(end.AddDate(0, 0, -increment)) == false; d = d.AddDate(0, 0, increment){
		sinceDateList = append(sinceDateList, d.Format("2006-01-02"))
	}
	untilStart := start.AddDate(0, 0, increment - 1)
	untilEnd := time.Now()
	for d := untilStart; d.After(untilEnd) == false; d = d.AddDate(0, 0, increment) {
		if(d.AddDate(0, 0, increment).After(untilEnd)){
			d = untilEnd
		}
		untilDateList = append(untilDateList, d.Format("2006-01-02"))
	}

	fmt.Println(start)
	fmt.Println(end)

	if Arguments.StartDate == "" && Arguments.EndDate == "" {
		fmt.Println("no dates")
		query := Arguments.Query
		for i := 0; i < len(sinceDateList); i = i + 1{
    		Core.Main(&(query), &(Arguments.Instance), &(Arguments.Format), &(Arguments.Name))
    	}
	} else if Arguments.StartDate != "" && Arguments.EndDate == "" {
		fmt.Println("start only")
		fmt.Println(len(sinceDateList))
		for i := 0; i < len(sinceDateList); i = i + 1{
			query := Arguments.Query + " since:" + sinceDateList[i]
			fmt.Println(query)
	    	Core.Main(&(query), &(Arguments.Instance), &(Arguments.Format), &(Arguments.Name))
	    }	
	} else if Arguments.StartDate == "" && Arguments.EndDate != "" {
		fmt.Println("end only")
		for i := 0; i < len(sinceDateList); i = i + 1{
			query := Arguments.Query + " until:" + untilDateList[i]	
	    	Core.Main(&(query), &(Arguments.Instance), &(Arguments.Format), &(Arguments.Name))
	    }
	} else if Arguments.StartDate != "" && Arguments.EndDate != "" {
		fmt.Println("both dates")
		for i := 0; i < len(sinceDateList); i = i + 1{
			query := Arguments.Query + " since:" + sinceDateList[i] + " until:" + untilDateList[i]	
	    	Core.Main(&(query), &(Arguments.Instance), &(Arguments.Format), &(Arguments.Name))
	    }	
	}
}	
