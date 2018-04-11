package twitch

import (
	"io"
	"io/ioutil"
	"net/http"
)

func TwitchRequest(method string, url string, body io.Reader, auth bool) (string, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Add("Client-ID", ClientID)
	if auth {
		req.Header.Add("Authorization", "OAuth "+Token)
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
