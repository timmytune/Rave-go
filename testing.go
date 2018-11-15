package main

import (
	"Rave-go/rave"
	"Rave-go/rave/transfer"
	"fmt"
)

type payload struct {
	PBFPubKey      string
	Amount         int    `json:"amount"`
	Txref          string `json:"txRef"`
	Customer_email string `json:"email"`
	Customer_phone string `json:"customer_phone"`
	Currency       string `json:"currency"`
	Cardno         string `json:"cardno"`
	Cvv            string `json:"cvv"`
	Expirymonth    int    `json:"expirymonth"`
	Expiryyear     int    `json:"expiryyear"`
	Pin            int    `json:"pin"`
}

func main() {

	h := rave.Rave{
		false,
		"FLWPUBK-ba0a57153f497c03bf34a9e296aa9439-X",
		"FLWSECK-327b3874ca8e75640a1198a1b75c0b0b-X",
	}
	fmt.Println(h)

	t := transfer.Transfer{
		h,
	}

	//paymentData := payload{
	//	PBFPubKey:h.GetPublicKey(),
	//	Amount:2000,
	//	Txref:"MC-11001993",
	//	Customer_email:"kwakujosh@gmail.com",
	//	Customer_phone:"09093146022",
	//	Currency:"NGN",
	//	Cardno:"5399830243533732",
	//	Cvv:"144",
	//	Expirymonth:10,
	//	Expiryyear:21,
	//	Pin:2424,
	//}
	//p := transfer.SinglePaymentData{
	//	AccountBank: "044",
	//	AccountNumber: "0690000044",
	//	Amount: 500,
	//	Narration: "New transfer",
	//	Currency: "NGN",
	//	Reference: "mk-90283-",
	//
	//}

	//q := transfer.BulkPaymentData{
	//
	//		SecKey:"FLWSECK-0b1d6669cf375a6208db541a1d59adbb-X",
	//		Title:"May Staff Salary",
	//		BulkData: []map[string]string{
	//			{
	//				"Bank":"044",
	//				"Account Number": "0690000032",
	//				"Amount":"500",
	//				"Currency":"NGN",
	//				"Narration":"Bulk transfer 1",
	//				"reference": "mk-82973029",
	//			},
	//			{
	//				"Bank":"044",
	//				"Account Number": "0690000034",
	//				"Amount":"500",
	//				"Currency":"NGN",
	//				"Narration":"Bulk transfer 2",
	//				"reference": "mk-283874750",
	//			},
	//		},
	//}

	a := transfer.AccountResolveData{

		RecipientAccount: "0690000034",
		DestBankCode:     "044",
		PublicKey:        "FLWPUBK-4e9d4e37974a61157ce8ca4f43c84936-X",
	}
	err, response := t.ResolveAccount(a)
	if err != nil {
		panic(err)
	}
	fmt.Println("response", response)

	//p, err := json.Marshal(paymentData)
	//if err != nil {
	//	log.Fatal("Could not create JSON of data")
	//}
	//fmt.Println(string(p))
	////h.GetKey(h.SecretKey)
	//e := h.Encrypt(string(p))
	//fmt.Println(e)
}
