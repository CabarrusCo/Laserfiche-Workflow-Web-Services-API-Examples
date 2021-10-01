package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	httpntlm "github.com/vadimi/go-http-ntlm/v2"
)

type StartWFRequest struct {
	ParameterCollection []Parameter `json:"ParameterCollection"`
}

type Parameter struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type WFResponse struct {
	Fault      Fault  `json:"fault"`
	InstanceID string `json:"instanceId"`
}

type Fault struct {
	Status     int    `json:"Status"`
	Detail     string `json:"Detail"`
	DetailCode int    `json:"DetailCode"`
}

func main() {
	swfr := StartWFRequest{
		ParameterCollection: []Parameter{
			{
				Name:  "Message",
				Value: "Hello World",
			},
		},
	}

	wfJSON, err := json.Marshal(swfr)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := http.Client{
		Transport: &httpntlm.NtlmTransport{
			Domain:   "XXX",
			User:     "XXXXXXXX",
			Password: "XXXXXX",
			// Configure RoundTripper if necessary, otherwise DefaultTransport is used
			RoundTripper: &http.Transport{
				// provide tls config
				TLSClientConfig: &tls.Config{},
				// other properties RoundTripper, see http.DefaultTransport
			},
		},
	}

	req, err := http.NewRequestWithContext(context.Background(), "POST", "workflowURLHere", bytes.NewBuffer(wfJSON))
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Wanted status code 200, got %v instead", resp.StatusCode)
		return
	}

	wfResponse := WFResponse{}

	err = json.NewDecoder(resp.Body).Decode(&wfResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	if wfResponse.Fault.Status != 0 {
		fmt.Printf("Fault encountered while starting workflow %v", wfResponse.Fault.Detail)
		return
	}

	fmt.Printf("Workflow started successfully with Instance ID %v", wfResponse.InstanceID)
}