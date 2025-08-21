package instruments

import (
	"encoding/json"
	"fmt"
	"ldv/tinvest"
	"log"
)

type SecurityDesc struct {
	Figi                  string
	Ticker                string
	ClassCode             string
	ISIN                  string
	Lot                   int
	Currency              string
	BuyAvailableFlag      bool
	SellAvailableFlag     bool
	ShortEnableFlag       bool
	Name                  string
	Exchange              string
	IpoDate               string
	CountryOfRisk         string
	CountryOfRiskName     string
	Sector                string
	IssueSizePlan         string
	DivYieldFlag          bool
	ShareType             string
	ApiTradeAvailableFlag bool
	Uid                   string
	PositionUid           string
	AssetUid              string
	ForQualInvestorFlag   bool
	WeekendFlag           bool
	Brand                 tinvest.Brand
	MinPriceIncrement     tinvest.Amount
	InstrumentType        tinvest.SecurityType
}

type Security struct {
	Instrument SecurityDesc
}

func SecurityBy(token string, figi string, itype tinvest.SecurityType) Security {
	url_shares := "https://invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.InstrumentsService/ShareBy"
	url_bondes := "https://invest-public-api.tbank.ru/rest/tinkoff.public.invest.api.contract.v1.InstrumentsService/BondBy"
	url_currences := "https://invest-public-api.tbank.ru/rest/tinkoff.public.invest.api.contract.v1.InstrumentsService/CurrencyBy"
	url_etfs := "https://invest-public-api.tbank.ru/rest/tinkoff.public.invest.api.contract.v1.InstrumentsService/EtfBy"
	url_futuries := "https://invest-public-api.tbank.ru/rest/tinkoff.public.invest.api.contract.v1.InstrumentsService/FutureBy"

	var url string

	switch itype {
	case tinvest.Share:
		url = url_shares
	case tinvest.Bond:
		url = url_bondes
	case tinvest.Currency:
		url = url_currences
	case tinvest.Etf:
		url = url_etfs
	case tinvest.Future:
		url = url_futuries
	default:
		log.Fatal("SecurityBy() Unknown type of security ")
	}

	var sec Security

	payload := fmt.Sprintf(`{"idType":"%s", "id":"%s"}`, "INSTRUMENT_ID_TYPE_FIGI", figi)
	data := tinvest.GetAPIRequest(url, token, payload)
	err := json.Unmarshal([]byte(data), &sec)
	if err != nil {
		log.Fatal(err)
	}
	return sec
}
