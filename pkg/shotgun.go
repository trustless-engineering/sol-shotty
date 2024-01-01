package pkg

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type SuccessResponse struct {
	Result   *http.Response `json:"result"`
	Endpoint string         `json:"endpoint"`
	RTT      int            `json:"rtt"`
	Body     []byte         `json:"-"`
}

func Shotgun(endpoints []string, mainRequest *http.Request) (SuccessResponse, error) {
	successCh := make(chan SuccessResponse)

	defer close(successCh)

	for _, endpoint := range endpoints {
		go makeRequest(endpoint, mainRequest, successCh)
	}

	return <-successCh, nil
}

func makeRequest(endpoint string, req *http.Request, successCh chan SuccessResponse) {
	// Create a new request with the same method and body as the original request
	newReq, err := http.NewRequest(req.Method, endpoint, req.Body)
	if err != nil {
		return
	}

	// Copy headers from the original request to the new request
	newReq.Header = make(http.Header)
	for key, values := range req.Header {
		newReq.Header[key] = values
	}

	// Send request
	client := &http.Client{
		Timeout: 30 * time.Second, //TODO: Make this configurable
	}

	startTime := time.Now()
	resp, err := client.Do(newReq)
	endTime := time.Now()

	if err != nil {
		return
	}

	defer resp.Body.Close()

	// Check for a successful response (status code 2xx)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}

		var jsonData map[string]interface{}
		err = json.Unmarshal(body, &jsonData)
		if err != nil {
			return
		}

		// Reject responses with an "error" key
		if _, ok := jsonData["error"]; ok {
			return
		}

		successCh <- SuccessResponse{
			Result:   resp,
			Endpoint: endpoint,
			RTT:      int(endTime.Sub(startTime).Milliseconds()),
			Body:     body,
		}

		return
	}
}
