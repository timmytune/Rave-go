package payment

import (
	"Rave-go/rave"
	"Rave-go/rave/helper"
	"go/types"
	// "fmt"
)


type CardCharge interface {
	ChargeCard(data CardChargeData) (error error, response map[string]interface{})
}

type CardValidate interface {
	ValidateCard(data CardValidateData) (error error, response map[string]interface{})
}

type CardVerify interface {
	VerifyCard(data CardVerifyData) (error error, response map[string]interface{})
}

type CardInterface interface {
	CardCharge
	CardValidate
	CardVerify
}

type CardChargeData struct {
	Cardno             string         `json:"cardno"`
	Cvv                string         `json:"cvv"`
	Expirymonth        string         `json:"expirymonth"`
	Expiryyear         string         `json:"expiryyear"`
	Pin                string         `json:"pin"`
	Amount             int            `json:"amount"`
	Currency           string         `json:"currency"`
	CustomerPhone      string         `json:"customer_phone"`
	Firstname          string         `json:"firstname"`
	Lastname           string         `json:"lastname"`
	Email              string         `json:"email"`
	Ip                 string         `json:"IP"`
	Txref		       string	      `json:"txRef"`
	RedirectUrl        string         `json:"redirect_url"`
	Subaccounts        types.Slice    `json:"subaccounts"`
	DeviceFingerprint  string         `json:"device_fingerprint"`
	Meta               types.Slice    `json:"meta"`
	SuggestedAuth      string         `json:"suggested_auth"`
	BillingZip         string          `json:"billingzip"`
	BillingCity        string           `json:"billingcity"`
	BillingAddress      string          `json:"billingaddress"`
	BillingState       string           `json:"billingstate"`
	BillingCountry      string          `json:"billingcountry"`

}

type CardValidateData struct {
	Reference	   string	      `json:"transaction_reference"`
	Otp		       string	      `json:"otp"`
	PublicKey      string         `json:"PBFPubKey"`
}

type CardVerifyData struct {
	Reference	   string	      `json:"txref"`
	Amount	       int	          `json:"amount"`
	Currency       string         `json:"currency"`
	SecretKey      string         `json:"SECKEY"`
}

type Card struct {
	rave.Rave
}

func (c Card) ChargeCard(data CardChargeData) (error error, response map[string]interface{}) {
	chargeJSON := helper.MapToJSON(data)
	encryptedChargeData := c.Encrypt(string(chargeJSON[:]))
	queryParam := map[string]interface{}{
        "PBFPubKey": c.GetPublicKey(),
        "client": encryptedChargeData,
        "alg": "3DES-24",
    }
	
	url := c.GetBaseURL() + c.GetEndpoint("card", "charge")
	err, response := helper.MakePostRequest(queryParam, url)
	if err != nil {
		return err, noresponse
	}
	suggestedAuth := response["data"].(map[string]interface{})["suggested_auth"]
	if (suggestedAuth == "PIN") {
		data.SuggestedAuth = "PIN"
		chargeJSON = helper.MapToJSON(data)
		encryptedChargeData = c.Encrypt(string(chargeJSON[:]))
		queryParam = map[string]interface{}{
			"PBFPubKey": c.GetPublicKey(),
			"client": encryptedChargeData,
			"alg": "3DES-24",
		}
		err, response = helper.MakePostRequest(queryParam, url)
		if err != nil {
			return err, noresponse
		}
	} else if (suggestedAuth == "AVS_VBVSECURECODE") {
		data.SuggestedAuth = "AVS_VBVSECURECODE"
		chargeJSON = helper.MapToJSON(data)
		encryptedChargeData = c.Encrypt(string(chargeJSON[:]))
		queryParam = map[string]interface{}{
			"PBFPubKey": c.GetPublicKey(),
			"client": encryptedChargeData,
			"alg": "3DES-24",
		}
		err, response = helper.MakePostRequest(queryParam, url)
		if err != nil {
			return err, noresponse
		}

	}

	return nil, response

}

func (c Card) ValidateCard(data CardValidateData) (error error, response map[string]interface{}) {
	data.PublicKey = c.GetPublicKey()
	url := c.GetBaseURL() + c.GetEndpoint("card", "validate")
	err, response := helper.MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response

}

func (c Card) VerifyCard(data CardVerifyData) (error error, response map[string]interface{}) {
	data.SecretKey = c.GetSecretKey()
	url := c.GetBaseURL() + c.GetEndpoint("card", "verify")
	err, response := helper.MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	
	return nil, response

}