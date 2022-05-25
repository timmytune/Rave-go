package rave

import (
	"testing"
)

func TestChargeCard(t *testing.T) {
	var tests = []CardChargeData{
		{
			Amount:        30,
			Txref:         "MC-11001993",
			Email:         "kwaku@gmail.com",
			CustomerPhone: "08123456789",
			Currency:      "NGN",
			Cardno:        "5399838383838381",
			Cvv:           "470",
			Expirymonth:   "10",
			Expiryyear:    "22",
			Pin:           "3310",
		},
	}

	for _, test := range tests {
		error, response := Card{
			r,
		}.ChargeCard(test)
		if error != nil {
			t.Fatalf("Card Charge failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Card Charge status: %v, Details: %v", response["status"], response)
		}
	}
}

// func TestValidateCard(t *testing.T) {
// 	var tests = []CardValidateData {
// 		{
// 			Otp:"12345",
// 			Reference: "FLW-MOCK-e80776317e0dc061dd7f04662f308e58",
// 			// resp["data"].(map[string]interface{})["flwRef"].(string),
// 		},

// 	}

// 	for _, test := range tests {
// 		error, response := Card{
// 			r,
// 		}.ValidateCard(test)
// 		if error != nil{
// 			t.Fatalf("Validate Charge failed with error %v",error)
// 		}
// 		if response["status"] != "success"{
// 			t.Fatalf("Validate Charge status: %v, Details: %v",response["status"], response)
// 		}
// 	}
// }

func TestVerifyCard(t *testing.T) {
	var tests = []CardVerifyData{
		{
			Amount:    30,
			Currency:  "NGN",
			Reference: "MC-11001993",
		},
	}

	for _, test := range tests {
		error, response := Card{
			r,
		}.VerifyCard(test)
		if error != nil {
			t.Fatalf("Verify Charge failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Verify Charge status: %v, Details: %v", response["status"], response)
		}
	}
}

// func TestChargePreauth(t *testing.T) {
// 	var tests = []CardChargeData  {
// 		{
// 			Amount:10,
// 			Txref:"MC-11001993",
// 			Email:"kwaku@gmail.com",
// 			CustomerPhone:"08123456789",
// 			Currency:"NGN",
// 			Cardno:"5399838383838381",
// 			Cvv:"470",
// 			Expirymonth:"10",
// 			Expiryyear:"22",
// 			Pin: "3310",
// 		},

// 	}

// 	for _, test := range tests {
// 		error, response := Card{
// 			r,
// 		}.ChargePreauth(test)
// 		if error != nil{
// 			t.Fatalf("Preauth Capture failed with error %v",error)
// 		}
// 		if response["status"] != "success"{
// 			t.Fatalf("Preauth Capture status: %v, Details: %v",response["status"], response)
// 		}
// 	}
// }
// func TestCapturePreauth(t *testing.T) {
// 	var tests = []PreauthCaptureData  {
// 		{
// 			Amount:100,
// 			Flwref:"FLW-MOCK-e80776317e0dc061dd7f04662f308e58",
// 		},

// 	}

// 	for _, test := range tests {
// 		error, response := Card{
// 			r,
// 		}.CapturePreauth(test)
// 		if error != nil{
// 			t.Fatalf("Preauth Capture failed with error %v",error)
// 		}
// 		if response["status"] != "success"{
// 			t.Fatalf("Preauth Capture status: %v, Details: %v",response["status"], response)
// 		}
// 	}
// }
