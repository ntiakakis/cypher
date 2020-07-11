package internal

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Fetch ... Fetches stuff
func Fetch(domain, token string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://"+domain+"/users/fetch", nil)
	req.Header.Add("Authorization", string(token))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}
