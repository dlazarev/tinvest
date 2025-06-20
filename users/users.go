package users

import (
	"encoding/json"
	"ldv/tinvest"
	"log"
	"time"
)

type Account struct {
	Id          string
	Type        string
	Name        string
	Status      string
	OpenedDate  time.Time
	ClosedDate  time.Time
	AccessLevel string `json:"accesslevel"`
}

type UserInfo struct {
	UserId               string
	Tariff               string
	QualStatus           bool
	PremStatus           bool
	RiskLevelCode        string
	QualifiedForWorkWith []string
}

type AccountsData struct {
	Accounts []Account `json:"accounts"`
}

type unaryLimit struct {
	LimitPerMinute int
	Methods        []string
}

type streamLimit struct {
	Streams []string
	Limit   int
	Open    int
}

type UserTariff struct {
	UnaryLimits  []unaryLimit
	StreamLimits []streamLimit
}

func GetInfo(token string) UserInfo {
	var userinfo UserInfo

	url := "https://invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.UsersService/GetInfo"
	datauser := tinvest.GetAPIRequest(url, token, `{}`)
	err := json.Unmarshal([]byte(datauser), &userinfo)
	if err != nil {
		log.Fatal(err)
	}
	return userinfo
}

func GetAccounts(token string) AccountsData {
	var accInfo AccountsData

	url := "https://invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.UsersService/GetAccounts"
	accdata := tinvest.GetAPIRequest(url, token, `{"status":"ACCOUNT_STATUS_UNSPECIFIED"}`)
	err := json.Unmarshal([]byte(accdata), &accInfo)
	if err != nil {
		log.Fatal(err)
	}

	return accInfo
}

func GetUserTariff(token string) UserTariff {
	var tariff UserTariff

	url := "https://invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.UsersService/GetUserTariff"
	data := tinvest.GetAPIRequest(url, token, `{}`)
	err := json.Unmarshal([]byte(data), &tariff)
	if err != nil {
		log.Fatal(err)
	}

	return tariff
}
