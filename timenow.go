package timenow

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

const (
	apiEndpoint string = "http://worldclockapi.com/api/json/utc/now"
	timeFormat  string = "2006-01-02T15:04Z"
)

type Timenow struct {
	HttpClient *http.Client
}

type timeResponse struct {
	Now string `json:"currentDateTime,omitempty"`
}

func New(httpClient *http.Client) *Timenow {
	hc := httpClient

	if hc == nil {
		hc = &http.Client{Timeout: 10 * time.Second}
	}

	return &Timenow{
		HttpClient: hc,
	}
}

func (t *Timenow) Execute() (string, error) {
	resp, err := t.HttpClient.Get(apiEndpoint)

	if err != nil {
		err = errors.Wrap(err, "failed to contact API")
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.Wrap(err, fmt.Sprintf("non successful API response (%d)", resp.StatusCode))
		return "", err
	}

	tr := timeResponse{}
	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		err = errors.Wrap(err, "cannot decode server response")
		return "", err
	}

	timeNow, err := time.Parse(timeFormat, tr.Now)
	if err != nil {
		err = errors.Wrap(err, "cannot parse time")
		return "", err
	}

	now := timeNow.Format(time.UnixDate)

	return now, nil
}
