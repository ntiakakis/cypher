package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(domain, username, password string) string {
	_byte, err := json.Marshal(map[string]string{"username": username, "password": password})

	if err != nil {
		log.Fatalln(err)
		return ""
	}

	resp, err := http.Post("http://"+domain+"/auth/generate", "application/json", bytes.NewBuffer(_byte))

	if err != nil {
		log.Fatalln(err)
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
		return ""
	}

	return string(body)
}
