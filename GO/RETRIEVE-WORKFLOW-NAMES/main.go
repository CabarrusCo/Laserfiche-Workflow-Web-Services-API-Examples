package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Azure/go-ntlmssp"
)

type workflow struct {
	Name string `json:"name"`
}

func main() {
	url, user, password := "URL", "XXXXXX", "XXXXXXX"

	client := &http.Client{
		Transport: ntlmssp.Negotiator{
			RoundTripper: &http.Transport{},
		},
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Println("Excepted 200, Got: ", res.StatusCode)
		return
	}

	workflows := []workflow{}
	err = json.NewDecoder(res.Body).Decode(&workflows)
	if err != nil {
		log.Println(err)
		return
	}

	for _, w := range workflows {
		log.Println("WORKFLOW NAME IS: ", w.Name)
	}
}
