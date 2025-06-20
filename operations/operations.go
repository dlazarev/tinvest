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
}

func GetPositions(token string, accountId string) string {
	url := "https://invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.OperationsService/GetPositions"
	payload := fmt.Sprintf(`{"accountId":"%s"}`, accountId)

	return tinvest.GetAPIRequest(url, token, payload)
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
