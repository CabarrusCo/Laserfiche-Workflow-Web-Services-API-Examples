package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Azure/go-ntlmssp"
)

type Parameter struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}

type workflowData struct {
	ParameterCollection []Parameter
}

func main() {

	workflowInfo := workflowData{
		ParameterCollection: []Parameter{
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

	url, user, password := "XXXXXXX", "XXXXXXXX", "XXXXXX"

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	//You can decode to JSON if you want
	log.Println(string(body))

}
