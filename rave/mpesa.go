package rave

import ( 
	// "fmt"
	"go/types"

)
//Sub Interface

type MpesaMobilemoneyCharge interface{
 MpesaMobilemoneyCharge (data MpesaMobileChargeData ) (error error, response map[string]interface{})
}

//Main Interface

type Mpesamobilemoney interface { 
	MpesaMobilemoneyCharge 
}
type MpesaMobileChargeData struct {
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
	Narration string `json:"narration"`
	Txref string `json:"txRef"`
	Meta  types.Slice    `json:"meta"`
	IsMpesa int `json:"is_mpesa"`
	IsMpesaLipa int `json:"is_mpesa_lipa"`
	DeviceFingerprint string `json:"device_fingerprint"`
	
}

type MpesaMobilemoney struct{
	Rave
}

func (m MpesaMobilemoney) SetupCharge(data MpesaMobileChargeData) map[string]interface{} {
	chargeJSON := MapToJSON(data)
	encryptedChargeData := m.Encrypt(string(chargeJSON[:]))
	queryParam := map[string]interface{}{
        "PBFPubKey": m.GetPublicKey(),
        "client": encryptedChargeData,
        "alg": "3DES-24",
    }
	return queryParam
}

func (m MpesaMobilemoney) MpesaMobilemoneyCharge (data MpesaMobileChargeData ) (error error, response map[string]interface{}) {

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

