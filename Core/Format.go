package Core

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"os"
)

func FormatTweets(format string, tweets []Tweet, Name *string) {
	if format == "json" {
		FormatTweetsJSON(tweets, Name)
	} else {
		FormatTweetsCSV(tweets, Name)
	}
}

func FormatTweetsCSV(tweets []Tweet, Name *string) {
	nameValue := *Name
	file, err := os.OpenFile(nameValue + ".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}

	var b []byte

	//writer to print to console
	buf := bytes.NewBuffer(b)
	w := csv.NewWriter(buf)

	//writer that exports to CSV
	csvW := csv.NewWriter(file)


	for _, tweet := range tweets {

		attachments := make([]string, len(tweet.Attachments))
		for i, att := range tweet.Attachments {
			attachments[i] = *att.URL
		}

		row := []string{
			tweet.ID,
			tweet.URL,
			tweet.Timestamp,
			tweet.Username,
			tweet.Fullname,
			tweet.Text,
			strings.Join(attachments, ","),
			fmt.Sprintf("%d", tweet.Stats.Replies),
			fmt.Sprintf("%d", tweet.Stats.Retweets),
			fmt.Sprintf("%d", tweet.Stats.Quotes),
			fmt.Sprintf("%d", tweet.Stats.Likes),
		}

		//write the string to the Writer
		if err := w.Write(row); err != nil {
			log.Fatalln("error writing row to print csv buffer:", err)
		}

		if err := csvW.Write(row); err != nil{
			log.Fatalln("error writing row to csv buffer:", err)
		}

		fmt.Print(row)
	}
	w.Flush()
	csvW.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(buf.Bytes()))
}

func FormatTweetsJSON(tweets []Tweet, Name *string) {
	for _, tweet := range tweets {
		tweetJSON, _ := json.Marshal(tweet)
		fmt.Println(string(tweetJSON))
	}
}