package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Azure/go-ntlmssp"
)

func main() {

	url, user, password := "XXXXXX", "XXXXXX", "XXXXXXX"

	client := &http.Client{
		Transport: ntlmssp.Negotiator{
			RoundTripper: &http.Transport{},
		},
	}

	jsonStr := []byte(`{"ParameterCollection":[{"Name":"Message", "Value":"Hello World!"}]}`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
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

	var wfResp interface{}
	err = json.NewDecoder(resp.Body).Decode(&wfResp)
	if err != nil {
		log.Println(err)
		return
	}

	wfRespData := wfResp.(map[string]interface{})

	log.Println("INSTANCE ID IS ", wfRespData["instanceId"])

}
