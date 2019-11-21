package rave

import ( 
	// "fmt"
	// "go/types"

)
//Sub Interface

type UgandaMobilemoneyCharge interface{
 UgandaMobilemoneyCharge (data UgandaMobileChargeData ) (error error, response map[string]interface{})
}

//Main Interface

type Ugandamobilemoney interface { 
	UgandaMobilemoneyCharge 
}
type UgandaMobileChargeData struct {
	Pubkey string `json:"PBFPubKey"`
	Currency string `json:"currency"`
	PaymentType string `json:"payment_type"`
	Country string `json:"country"`
	Amount string `json:"amount"`
	Email string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Network string `json:"network"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	IP string `json:"IP"`
	Txref string `json:"txRef"`
	OrderRef string `json:"orderRef"`
	IsMobile int `json:"is_mobile_money_ug"`
	RedirectURL string `json:"redirect_url"`
	DeviceFingerprint string `json:"device_fingerprint"`
}

type UgandaMobilemoney struct{
	Rave
}

func (m UgandaMobilemoney) SetupCharge(data UgandaMobileChargeData) map[string]interface{} {
	chargeJSON := MapToJSON(data)
	encryptedChargeData := m.Encrypt(string(chargeJSON[:]))
	queryParam := map[string]interface{}{
        "PBFPubKey": m.GetPublicKey(),
        "client": encryptedChargeData,
        "alg": "3DES-24",
    }
	return queryParam
}

func (m UgandaMobilemoney) UgandaMobilemoneyCharge (data UgandaMobileChargeData ) (error error, response map[string]interface{}) {

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

