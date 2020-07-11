package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func Register(domain, username, password, email string) string {
	_byte, err := json.Marshal(map[string]string{"username": username, "password": password, "email": email})

	if err != nil {
		log.Fatalln(err)
		return ""
	}

	resp, err := http.Post("http://"+domain+"/users/register", "application/json", bytes.NewBuffer(_byte))

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
