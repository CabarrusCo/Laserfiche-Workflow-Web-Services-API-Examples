package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Azure/go-ntlmssp"
)

type workflowData struct {
	ParameterCollection []parameter
}

type parameter struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type workflowResponse struct {
	Fault      workflowFault `json:"fault"`
	InstanceID string        `json:"instanceId"`
}
type workflowFault struct {
	Status     int    `json:"Status"`
	Detail     string `json:"Detail"`
	DetailCode int    `json:"DetailCode"`
}

func main() {

	workflowInfo := workflowData{
		ParameterCollection: []parameter{
			{
				Name:  "Message",
				Value: "Hello World!",
			},
			{
				Name:  "MESSAGE TWO",
				Value: "HELLO WORLD AGAIN!",
			},
		},
	}

	wfJSON, err := json.Marshal(workflowInfo)
	if err != nil {
		log.Println(err)
		return
	}

	url, user, password := "XXXXXXX", "XXXXX", "XXXXXX"

	client := &http.Client{
		Transport: ntlmssp.Negotiator{
			RoundTripper: &http.Transport{},
		},
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(wfJSON))
	if err != nil {
		log.Println(err)
		return
	}

	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	wfResponse := workflowResponse{}
	err = json.NewDecoder(resp.Body).Decode(&wfResponse)
	if err != nil {
		log.Println("ERROR DECODING RESPONSE")
	}

	if wfResponse.Fault.Status != 0 {
		log.Println("PROBLEM STARTING WORKFLOW:: ", wfResponse.Fault.Detail)
		return
	}

	log.Println("WORKFLOW STARTED SUCCESSFULLY WITH INSTANCE ID:: ", wfResponse.InstanceID)

}
