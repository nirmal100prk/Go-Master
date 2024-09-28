package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// url := "https://is.gd/create.php?format=simple&url="
	// httpReq.SendHttpRequest()
	s, err := ShortenURL("https://gitlab.dt.ae/cerebrum/business-app/fms/mussad-service/-/merge_requests/52")
	fmt.Println(s)
	fmt.Println(err)
}

func ShortenURL(longURL string) (string, error) {
	// Prepare the API URL
	apiURL := "https://is.gd/create.php?format=simple&url=" + url.QueryEscape(longURL)

	// Send the GET request to the API
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to shorten URL: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	// Return the shortened URL as a string
	return string(body), nil
}
