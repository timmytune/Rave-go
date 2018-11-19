package payment

import (
	"Rave-go/rave"
	"Rave-go/rave/helper"
	"go/types"
)

var noresponse = map[string]interface{}{
	"": "",
}

type AccountCharge interface {
	ChargeAccount(data AccountChargeData) (error error, response map[string]interface{})
}

type AccountValidate interface {
	ValidateAccount(data AccountValidateData) (error error, response map[string]interface{})
}

type AccountVerify interface {
	VerifyAccount(data AccountVerifyData) (error error, response map[string]interface{})
}

type AccountInterface interface {
	AccountCharge
	AccountValidate
	AccountVerify
}

type AccountChargeData struct {
	Cardno             string         `json:"cardno"`
	Cvv                string         `json:"cvv"`
	Accountbank        string         `json:"accountbank"`
	Accountnumber         string         `json:"accountnumber"`
	Paymenttype            string      `json:"payment_type"`
	Amount             int            `json:"amount"`
	Currency           string         `json:"currency"`
	Country          string         `json:"country"`
	Bvn             string            `json:"bvn"`
	Passcode             string            `json:"passcode"`
	CustomerPhone      string         `json:"customer_phone"`
	Firstname          string         `json:"firstname"`
	Lastname           string         `json:"lastname"`
	Email              string         `json:"email"`
	IP                 string         `json:"IP"`
	Txref		       string	      `json:"txRef"`
	RedirectUrl        string         `json:"redirect_url"`
	Subaccounts        types.Slice    `json:"subaccounts"`
	DeviceFingerprint  string         `json:"device_fingerprint"`
	Meta               types.Slice    `json:"meta"`
	SuggestedAuth      string         `json:"suggested_auth"`
}

type AccountValidateData struct {
	PublicKey        string           `json:"PBFPubKey"`
	Reference   string           `json:"transactionreference"`
	Otp              string           `json:"otp"`
}

type AccountVerifyData struct {
	Reference	   string	      `json:"txref"`
	Amount	       int	          `json:"amount"`
	Currency       string         `json:"currency"`
	SecretKey      string         `json:"SECKEY"`
}

type Account struct {
	rave.Rave
}

func (a Account) ChargeAccount(data AccountChargeData) (error error, response map[string]interface{}) {
	chargeJSON := helper.MapToJSON(data)
	encryptedChargeData := a.Encrypt(string(chargeJSON[:]))
	queryParam := map[string]interface{}{
        "PBFPubKey": a.GetPublicKey(),
        "client": encryptedChargeData,
        "alg": "3DES-24",
    }
	
	url := a.GetBaseURL() + a.GetEndpoint("account", "charge")
	err, response := helper.MakePostRequest(queryParam, url)
	if err != nil {
		return err, noresponse
	}

	return nil, response

}

// Validates account charge using otp
func (a Account) ValidateAccount(data AccountValidateData) (error error, response map[string]interface{}) {
	data.PublicKey = a.GetPublicKey()
    url := a.GetBaseURL() + a.GetEndpoint("account", "validate")
    err, response := helper.MakePostRequest(data, url)
    if err != nil {
        return err, noresponse
    }

    return nil, response
}

func (a Account) VerifyAccount(data AccountVerifyData) (error error, response map[string]interface{}) {
	data.SecretKey = a.GetSecretKey()
    url := a.GetBaseURL() + a.GetEndpoint("account", "verify")
    err, response := helper.MakePostRequest(data, url)
    if err != nil {
        return err, noresponse
    }

    return nil, response
}