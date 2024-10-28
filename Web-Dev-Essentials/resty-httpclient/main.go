package httpReq

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
)

// APIResponse holds the response from an API call
type APIResponse struct {
	StatusCode int
	Body       []byte
	Headers    map[string][]string
}

// SendHttpRequest makes a REST API call using Resty with the provided method, URL, headers, and body.
// It returns the response and any error encountered.
func SendHttpRequest(ctx context.Context, method, url string, headers map[string]string, body interface{}) (*APIResponse, error) {
	// Create a new Resty client
	client := resty.New()

	// Set a timeout for the request
	client.SetTimeout(10 * time.Second)

	// Initialize a new request
	req := client.R().SetContext(ctx).SetHeaders(headers)

	// Set the request body if provided
	if body != nil {
		req.SetBody(body)
	}

	// Execute the request based on the HTTP method
	var resp *resty.Response
	var err error

	switch method {
	case http.MethodGet:
		resp, err = req.Get(url)
	case http.MethodPost:
		resp, err = req.Post(url)
	case http.MethodPut:
		resp, err = req.Put(url)
	case http.MethodDelete:
		resp, err = req.Delete(url)
	default:
		return nil, fmt.Errorf("unsupported HTTP method: %s", method)
	}

	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}

	// Construct and return the APIResponse
	return &APIResponse{
		StatusCode: resp.StatusCode(),
		Body:       resp.Body(),
		Headers:    resp.Header(),
	}, nil
}

// func main() {
// 	// Example usage
// 	ctx := context.Background()
// 	url := "https://jsonplaceholder.typicode.com/posts"
// 	method := "GET"
// 	headers := map[string]string{
// 		"Content-Type": "application/json",
// 	}

// 	// Call the reusable function
// 	response, err := SendHttpRequest(ctx, method, url, headers, nil)
// 	if err != nil {
// 		fmt.Printf("Error making API call: %v\n", err)
// 		return
// 	}

// 	// Print the response
// 	fmt.Printf("Status Code: %d\n", response.StatusCode)
// 	fmt.Printf("Response Body: %s\n", string(response.Body))
// 	fmt.Printf("Response Headers: %v\n", response.Headers)
// }
