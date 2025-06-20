package tinvest

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func GetAPIRequest(url string, token string, payload_str string) string {
	method := "POST"
	payload := strings.NewReader(payload_str)

	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
