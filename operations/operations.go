package operations

import (
	"encoding/json"
	"fmt"
	"ldv/tinvest"
	"ldv/tinvest/instruments"

	//	"ldv/tinvest/instruments"
	"log"
	"strconv"
)

//type Sum interface {
//	Sum() float64
//}

type Portfolio struct {
	TotalAmountPortfolio tinvest.Amount
	DailyYield           tinvest.Amount
	DailyYieldRelative   tinvest.Amount
	TotalAmountBonds     tinvest.Amount // Облигации
	TotalAmountFutures   tinvest.Amount // Фьючерсы
	TotalAmountSp        tinvest.Amount
	TotalAmountEtf       tinvest.Amount // Фонды
	TotalAmountShares    tinvest.Amount // Акции
}

type IntString int

func (i *IntString) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		var s string
		if err := json.Unmarshal(data, &s); err != nil {
			return err
		}
		v, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		*i = IntString(v)
		return nil
	}

	var v int
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	*i = IntString(v)
	return nil
}

type Security struct {
	Figi            string
	Blocked         IntString `json:"blocked"`
	Balance         IntString `json:"balance"`
	PositionUid     string
	Ticker          string
	ExchangeBlocked bool
	InstrumentType  tinvest.SecurityType
	InstrumentDesc  instruments.SecurityDesc
}

type Positions struct {
	Money      []tinvest.Amount
	Blocked    []tinvest.Amount
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

	for i := range positions.Securities {
		secDesc := instruments.SecurityBy(token, positions.Securities[i].Figi, positions.Securities[i].InstrumentType)
		positions.Securities[i].InstrumentDesc = secDesc.Instrument
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
