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

	var url = "https://invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.InstrumentsService/"

	switch itype {
	case tinvest.Share:
		url += "ShareBy"
	case tinvest.Bond:
		url += "BondBy"
	case tinvest.Currency:
		url += "CurrencyBy"
	case tinvest.Etf:
		url += "EtfBy"
	case tinvest.Future:
		url += "FutureBy"
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
