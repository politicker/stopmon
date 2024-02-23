package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/errornil/mta/v2"
)

func main() {
	client, err := mta.NewFeedsClient(
		&http.Client{
			Timeout: 30 * time.Second,
		},
		os.Getenv("MTA_API_KEY"),
		"github.com/politicker/stopmon",
	)
	if err != nil {
		panic(err)
	}

	resp, err := client.GetFeedMessage(mta.FeedL)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}
