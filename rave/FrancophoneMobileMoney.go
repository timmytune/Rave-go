package rave

import ( 
	// "fmt"
	// "go/types"

)
//Sub Interface

type FrancoMobilemoneyCharge interface{
	FrancoMobilemoneyCharge (data FrancoMobileChargeData ) (error error, response map[string]interface{})
}

//Main Interface

type Francomobilemoney interface { 
	FrancoMobilemoneyCharge 
}
type FrancoMobileChargeData struct {
	Pubkey string `json:"PBFPubKey"`
	Currency string `json:"currency"`
	PaymentType string `json:"payment_type"`
	Country string `json:"country"`
	Amount string `json:"amount"`
	Email string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	IP string `json:"IP"`
	Txref string `json:"txRef"`
	OrderRef string `json:"orderRef"`
	IsMobile int `json:"is_mobile_money_franco"`
	DeviceFingerprint string `json:"device_fingerprint"`
}

type FrancoMobilemoney struct{
	Rave
}

func (m FrancoMobilemoney) SetupCharge(data FrancoMobileChargeData) map[string]interface{} {
	chargeJSON := MapToJSON(data)
	encryptedChargeData := m.Encrypt(string(chargeJSON[:]))
	queryParam := map[string]interface{}{
        "PBFPubKey": m.GetPublicKey(),
        "client": encryptedChargeData,
        "alg": "3DES-24",
    }
	return queryParam
}

func (m FrancoMobilemoney) FrancoMobilemoneyCharge (data FrancoMobileChargeData ) (error error, response map[string]interface{}) {

	var url string
	url = m.GetBaseURL() + m.GetEndpoint("mobilemoney", "charge")
	if (data.Txref == "") {
		data.Txref = GenerateRef()
	}
	postData := m.SetupCharge(data)

	err, response := MakePostRequest(postData, url)
	if err != nil {
		return err, noresponse
	}

	return	nil, response

	
}

