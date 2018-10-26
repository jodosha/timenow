package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	apiEndpoint string = "http://worldclockapi.com/api/json/utc/now"
	timeFormat  string = "2006-01-02T15:04Z"
)

type timeResponse struct {
	Now string `json:"currentDateTime,omitempty"`
}

func main() {
	resp, err := http.Get(apiEndpoint)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("cannot contact server (HTTP response %d)", resp.StatusCode))
		os.Exit(1)
	}

	tr := timeResponse{}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("cannot decode server response (%s)", err.Error()))
		os.Exit(1)
	}

	timeNow, err := time.Parse(timeFormat, tr.Now)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("cannot parse time (%s)", err.Error()))
		os.Exit(1)
	}

	now := timeNow.Format(time.UnixDate)

	fmt.Println(now)
}
