package twitch

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Request - request for twitch api
func Request(method string, url string, body io.Reader, auth bool, context bool) (string, error) {
	fmt.Println("Twitch Request: ", url)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Add("Client-ID", ClientID)
	if auth {
		req.Header.Add("Authorization", "OAuth "+Token)
	}
	if context {
		req.Header.Add("Content-Type", "application/json")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	rbody, _ := ioutil.ReadAll(resp.Body)

	return string(rbody), err
}
