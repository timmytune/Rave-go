package refund

import (
	"Rave-go/rave"
	"testing"


)

var r = rave.Rave{
	false,
	"FLWPUBK-f54d8d24292e377a71620bd82a8bb17c-X",
	"FLWSECK-a18ca169cb007a93db4479aff683a387-X",
}

func TestRefundTransaction(t *testing.T) {
	var tests = []RefundData {
		{
			Ref: "FLW-MOCK-476a260e67df43988a2ffeddf8e02cc2",
			Amount: 10,
		},
		
	}

	for _, test := range tests {
		error, response := Refund{
			r,
		}.RefundTransaction(test)
		if error != nil{
			t.Fatalf("Refund Transaction failed with error %v",error)
		}
		if response["status"] != "success"{
			t.Fatalf("Refund Transaction status: %v, Details: %v",response["status"], response)
		}
	}
}
