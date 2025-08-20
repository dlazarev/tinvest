package tinvest

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Amount struct {
	Currency string
	Units    string
	Nano     int
}

func (a Amount) Sum() float64 {
	sum, err := strconv.ParseFloat(a.Units, 32)
	if err != nil {
		log.Fatal(err)
	}
	sum += float64(a.Nano) / 10e8
	return sum
}

type SecurityType string

type Brand struct {
	LogoName      string
	LogoBaseColor string
	TextColor     string
}

const (
	Bond     SecurityType = "bond"
	Currency SecurityType = "currency"
	Etf      SecurityType = "etf"
	Future   SecurityType = "future"
	Option   SecurityType = "option"
	Share    SecurityType = "share"
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
