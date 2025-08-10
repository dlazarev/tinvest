package operations

import (
	"encoding/json"
	"fmt"
	"ldv/tinvest"
	"log"
	"strconv"
)

//type Sum interface {
//	Sum() float64
//}

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

type Portfolio struct {
	TotalAmountPortfolio Amount
	DailyYield           Amount
	DailyYieldRelative   Amount
	TotalAmountBonds     Amount // Облигации
	TotalAmountFutures   Amount // Фьючерсы
	TotalAmountSp        Amount
	TotalAmountEtf       Amount // Фонды
	TotalAmountShares    Amount // Акции
}

type SecurityType string

const (
	Bond     SecurityType = "bond"
	Currency SecurityType = "currency"
	Etf      SecurityType = "etf"
	Future   SecurityType = "future"
	Option   SecurityType = "option"
	Share    SecurityType = "share"
)

type Security struct {
	Figi            string
	Blocked         string
	Balance         string
	PositionUid     string
	Ticker          string
	ExchangeBlocked bool
	InstrumentType  string
}

type Positions struct {
	Money      []Amount
	Blocked    []Amount
	Securities []Security
}

func GetPositions(token string, accountId string) Positions {
	url := "https://invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.OperationsService/GetPositions"
	payload := fmt.Sprintf(`{"accountId":"%s"}`, accountId)

	var positions Positions

	data := tinvest.GetAPIRequest(url, token, payload)
	err := json.Unmarshal([]byte(data), &positions)
	if err != nil {
		log.Fatal(err)
	}
	return positions
}

func GetPortfolio(token string, accoundId string) Portfolio {
	var portfolio Portfolio

	url := "https://invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.OperationsService/GetPortfolio"
	payload := fmt.Sprintf(`{"accountId":"%s", "currency":"RUB"}`, accoundId)
	data := tinvest.GetAPIRequest(url, token, payload)

	err := json.Unmarshal([]byte(data), &portfolio)
	if err != nil {
		log.Fatal(err)
	}
	return portfolio
}
