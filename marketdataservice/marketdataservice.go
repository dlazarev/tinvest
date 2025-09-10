package marketdataservice

import (
	"encoding/json"
	"ldv/tinvest"
	"log"
	"strings"
)

type LastPrice struct {
	Price         tinvest.Amount
	Time          string
	InstrumentId  string
	LastPriceType string
}
type Prices struct {
	LastPrices []LastPrice
}

func GetLastPrices(token string, figies []string) Prices {
	var url = `https://invest-public-api.tbank.ru/rest/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastPrices`
	var lps Prices

	payload := "{ [" + strings.Join(figies, ", ") + "],}"
	data := tinvest.GetAPIRequest(url, token, payload)
	err := json.Unmarshal([]byte(data), &lps)
	if err != nil {
		log.Fatal(err)
	}
	return lps

}
