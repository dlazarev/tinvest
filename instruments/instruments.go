package instruments

import (
	"fmt"
	"ldv/tinvest"
	"ldv/tinvest/operations"
)

type SecurityDesc struct {
	InstrumentType operations.SecurityType
}

func SecurityBy(figi string, itype operations.SecurityType) SecurityDesc {
	url := "https://invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.InstrumentsService/ShareBy"
	payload := fmt.Sprintf(`{"idType":":"%s", "id":"%s"}`, "INSTRUMENT_ID_TYPE_FIGI", figi)
	data := tinvest.GetAPIRequest(url, token, payload)

}
