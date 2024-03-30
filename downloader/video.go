package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// VideoDownloader is a struct that represents a video downloader that can retrieve video download links from the given URL.
type VideoDownloader struct {
	apiKey  string
	apiHost string
	baseURL string
}

// NewVideDownloader creates a new instance of the VideoDownloader struct with provided apiKey.
func NewVideDownloader(apiKey string) *VideoDownloader {
	return &VideoDownloader{
		apiKey:  apiKey,
		apiHost: "all-media-downloader.p.rapidapi.com",
		baseURL: "https://all-media-downloader.p.rapidapi.com/rapid_download/download",
	}
}

// Download sends a POST request to the video downloader API and returns the response about Video.
func (d *VideoDownloader) Download(url string) (*VideoResponse, error) {
	// payload to send in the request
	payload := strings.NewReader("url=" + url)

	// create a new request
	req, err := http.NewRequest("POST", d.baseURL, payload)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// set the request headers
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("X-RapidAPI-Key", d.apiKey)
	req.Header.Add("X-RapidAPI-Host", d.apiHost)

	// send the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// parse the response body and to the struct
	var videoResponse VideoResponse
	if err = json.Unmarshal(body, &videoResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %v", err)
	}

	// return an error if the response status is failed
	if videoResponse.Status == ResponseStatusFailed {
		return nil, fmt.Errorf("error downloading video: %v", videoResponse.Message)
	}

	return &videoResponse, nil
}
