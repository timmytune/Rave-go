package rave

import (
	"Rave-go/rave/helper"
)

var noresponse = map[string]interface{}{
	"": "",
}

type RefundCharge interface {
	RefundTransaction(data RefundData) (error error, response map[string]interface{})
}

type RefundInterface interface {
	RefundCharge
}

type RefundData struct {
	Ref		       string	      `json:"ref"`
	Amount         int            `json:"amount"`
	SecretKey      string         `json:"seckey"`
}


func (r Rave) RefundTransaction(data RefundData) (error error, response map[string]interface{}) {
	data.SecretKey = r.GetSecretKey()
	url := r.GetBaseURL() + r.GetEndpoint("refund", "refund")
	err, response := helper.MakePostRequest(data, url)
	if err != nil {
		return err, noresponse
	}
	return nil, response

}
