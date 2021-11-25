//********************************************************************************************************************//
//
// Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
//
// This file is part of uptime-kuma-server-push.
// All code may be used. Feel free and maybe code something better.
//
// Author: Jonas Kwiedor (aka gowizzard)
//
//********************************************************************************************************************//

package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"net/http"
	"os"
	"time"
)

// UptimeKumaReturn is to decode the json return
type UptimeKumaReturn struct {
	Ok  bool   `json:"ok"`
	Msg string `json:"msg,omitempty"`
}

func main() {

	// Get environment & check the length
	url := os.Getenv("URL")
	msg := os.Getenv("MSG")
	cron := os.Getenv("CRON")

	// Define cronjobs
	s := gocron.NewScheduler(time.UTC)

	s.Cron(cron).Do(func() {

		// To get milliseconds
		start := time.Now()

		// Define client for request
		client := &http.Client{}

		// Define request & add url parameter
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatalln(err)
		}

		parameters := request.URL.Query()
		parameters.Add("msg", msg)
		parameters.Add("ping", fmt.Sprintf("%d", time.Since(start).Milliseconds()))
		request.URL.RawQuery = parameters.Encode()

		// Define response & send request
		response, err := client.Do(request)
		if err != nil {
			log.Fatalln(err)
		}

		// Decode json return
		var decode UptimeKumaReturn

		err = json.NewDecoder(response.Body).Decode(&decode)
		if err != nil {
			log.Fatalln(err)
		}

		// Check decode & log to console
		if decode.Ok {
			log.Println("The heartbeat was successfully sent to your system.")
		}

	})

	// Start scheduler
	s.StartBlocking()

}
