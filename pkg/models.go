package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func fetch(n int) (*ComicResult, error) {

	client := &http.Client{
		Timeout: 5 * time.Minute,
	}

	// concatenate strings to get url; ex: https://xkcd.com/571/info.0.json
	url := strings.Join([]string{Url, fmt.Sprintf("%d", n), "info.0.json"}, "/")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("http request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http err: %v", err)
	}

	var data ComicResult

	// error from web service, empty struct to avoid disruption of process
	if resp.StatusCode != http.StatusOK {
		data = ComicResult{}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, fmt.Errorf("json err: %v", err)
		}
	}

	resp.Body.Close()

	return &data, nil
}

func fetchTotalComic() (int, error) {

	client := &http.Client{
		Timeout: 5 * time.Minute,
	}

	// concatenate strings to get url; ex: https://xkcd.com/571/info.0.json
	url := strings.Join([]string{Url, "info.0.json"}, "/")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return -1, fmt.Errorf("http request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return -1, fmt.Errorf("http err: %v", err)
	}

	var data ComicResult

	// error from web service, empty struct to avoid disruption of process
	if resp.StatusCode != http.StatusOK {
		data = ComicResult{}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return -1, fmt.Errorf("json err: %v", err)
		}
	}

	resp.Body.Close()

	return data.Num, nil
}
