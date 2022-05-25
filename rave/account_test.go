package rave

import (
	"testing"
	// "fmt"
)

func TestChargeAccount(t *testing.T) {
	var tests = []AccountChargeData{
		{
			Accountbank:   "044",
			Accountnumber: "0690000031",
			Amount:        100,
			Country:       "NG",
			Currency:      "NGN",
			Email:         "ajb@yahoo.com",
			CustomerPhone: "08123456789",
			Firstname:     "Anjola",
			Lastname:      "Bassey",
			Paymenttype:   "account",
			IP:            "103.238.105.185",
			Txref:         "MXX-ASC-4578",
		},
	}

	for _, test := range tests {
		error, response := Account{
			r,
		}.ChargeAccount(test)
		if error != nil {
			t.Fatalf("Card Charge failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Card Charge status: %v, Details: %v", response["status"], response)
		}
	}
}

// func TestValidateAccount(t *testing.T) {
// 	var tests = []AccountValidateData {
// 		{
// 			Otp:"12345",
// 			Reference: "FLW-MOCK-e80776317e0dc061dd7f04662f308e58",
// 			// resp["data"].(map[string]interface{})["flwRef"].(string),
// 		},

// 	}

// 	for _, test := range tests {
// 		error, response := Account{
// 			r,
// 		}.ValidateAccount(test)
// 		if error != nil{
// 			t.Fatalf("Validate Charge failed with error %v",error)
// 		}
// 		if response["status"] != "success"{
// 			t.Fatalf("Validate Charge status: %v, Details: %v",response["status"], response)
// 		}
// 	}
// }

func TestVerifyAccount(t *testing.T) {
	var tests = []AccountVerifyData{
		{
			Amount:    100,
			Currency:  "NGN",
			Reference: "MXX-ASC-4578",
		},
	}

	for _, test := range tests {
		error, response := Account{
			r,
		}.VerifyAccount(test)
		if error != nil {
			t.Fatalf("Verify Charge failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Verify Charge status: %v, Details: %v", response["status"], response)
		}
	}
}
