package marketdataservice

import (
	"encoding/json"
	"ldv/tinvest"
	"log"
)

type LastPrice struct {
	Price         tinvest.Amount
	Time          string
	InstrumentUid string
	LastPriceType string
}
type Prices struct {
	LastPrices []LastPrice
}

func GetLastPrices(token string, instrumentId []string) Prices {
	var url = `https://invest-public-api.tbank.ru/rest/tinkoff.public.invest.api.contract.v1.MarketDataService/GetLastPrices`
	var lps Prices

	i, err := json.Marshal(instrumentId)
	if err != nil {
		log.Fatalf("Error in GetLastPrices(): %v", err)
	}

	payload := `{ "instrumentId":` + string(i) + `}`
	data := tinvest.GetAPIRequest(url, token, payload)
	err = json.Unmarshal([]byte(data), &lps)
	if err != nil {
		log.Fatal(err)
	}
	return lps

}
