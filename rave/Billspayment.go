package rave

// import (
// 	"time"
// )

type Flybuy interface {
	Flybuy(data FlyBuyData) (error error, response map[string]interface{})
}

// type FlyBuyBulk interface {
// 	FlyBuyBulk(data FlyBuyBulkData) (error error, response map[string]interface{})
// }

// type FlyHistory interface {
// 	Get(data FlyHistoryData) (error error, response map[string]interface{})
// }

type servicepayloaddata struct {
	Country       string `json:"Country"`
	Amount        int    `json:"Amount"`
	CustomerId    string `json:"CustomerId"`
	RecurringType int    `json:"RecurringType"`
	IsAirtime     bool   `json:"IsAirtime"`
	BillerName    string `json:"BillerName"`
	Reference     string `json:"Reference"`
}

type FlyBuyData struct {
	Service        string `json:"service"`
	ServiceMethod  string `json:"service_method"`
	ServiceVersion string `json:"service_version"`
	ServiceChannel string `json:"service_channel`
	ServicePayload servicepayloaddata
	Seckey         string `json:"secret_key"`
}

// type FlyBuyBulkData struct {
// 	Page   string `json:"page"`
// 	Seckey string `json:"seckey"`
// }

// type FlyHistoryData struct {
// 	Id     int32  `json:"page"`
// 	Seckey string `json:"seckey"`
// }

type Billpayment struct {
	Rave
}

func (b Billpayment) Flybuy(data FlyBuyData) (error error, response map[string]interface{}) {
	data.Seckey = b.GetSecretKey()
	url := b.GetBaseURL() + b.GetEndpoint("Billspayments", "flybuy")
	err, response := MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}

// func (v Virtualcards) List(data ListData) (error error, response map[string]interface{}) {
// 	data.Seckey = v.GetSecretKey()
// 	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "list")
// 	err, response := MakePostRequest(data, url)
// 	if err != nil {
// 		return err, noresponse
// 	}
// 	return nil, response
// }

// func (v Virtualcards) Get(data GetData) (error error, response map[string]interface{}) {
// 	data.Seckey = v.GetSecretKey()
// 	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "get")
// 	err, response := MakePostRequest(data, url)
// 	if err != nil {
// 		return err, noresponse
// 	}
// 	return nil, response
// }

// func (v Virtualcards) Fund(data FundData) (error error, response map[string]interface{}) {
// 	data.Seckey = v.GetSecretKey()
// 	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "fund")
// 	err, response := MakePostRequest(data, url)
// 	if err != nil {
// 		return err, noresponse
// 	}
// 	return nil, response
// }

// func (v Virtualcards) Withdraw(data WithdrawData) (error error, response map[string]interface{}) {
// 	data.Seckey = v.GetSecretKey()
// 	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "withdraw")
// 	err, response := MakePostRequest(data, url)
// 	if err != nil {
// 		return err, noresponse
// 	}
// 	return nil, response
// }

// func (v Virtualcards) Freeze(data FreezeData) (error error, response map[string]interface{}) {
// 	data.Seckey = v.GetSecretKey()
// 	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "freeze")
// 	url += data.CardId
// 	url += "/status/"
// 	url += data.StatusAction
// 	err, response := MakePostRequest(data, url)
// 	if err != nil {
// 		return err, noresponse
// 	}
// 	return nil, response
// }

// func (v Virtualcards) Fetch(data FetchData) (error error, response map[string]interface{}) {
// 	data.Seckey = v.GetSecretKey()
// 	url := v.GetBaseURL() + v.GetEndpoint("virtualcard", "fetch")
// 	err, response := MakePostRequest(data, url)
// 	if err != nil {
// 		return err, noresponse
// 	}
// 	return nil, response
// }
