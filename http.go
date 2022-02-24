package getui

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func sendPost(url string, authToken string, bodyByte []byte) (string, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	body := bytes.NewBuffer(bodyByte)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}

	req.Header.Add("authtoken", authToken)
	req.Header.Add("Charset", "UTF-8")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("request getui failed.", resp)
		return "", err
	}

	return string(result), nil
}
