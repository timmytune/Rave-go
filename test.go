package main

import (
	"fmt"

	"./rave"
	// "time"
	// "strconv"
)

func main() {
	var r = rave.Rave{
		false,
		"FLWPUBK_TEST-9070a8d235099f5ac7db26a0356e92b3-X",
		"FLWSECK_TEST-9e54889bc26206218e96152ce4f477f9-X",
	}
	var Beneficiary = rave.Billpayment{
		r,
	}
	// servicepayloaddata:=rave. ServicepayLoadData{
	// 	Country:       "NG",
	// 	Amount:        500,
	// 	CustomerId:    "+23490803840303",
	// 	RecurringType: 0,
	// 	IsAirtime:     true,
	// 	BillerName:    "AIRTIME",
	// 	Reference:     "9300049404444",
	// }
	payload := rave.FlyBuyData{
		Seckey:         "FLWSECK_TEST-9e54889bc26206218e96152ce4f477f9-X",
		Service:        "fly_recurring",
		ServiceMethod:  "get",
		ServiceVersion: "v1",
		ServiceChannel: "rave",
		// ServicePayload: servicepayloaddata,
	}

	fmt.Println(payload)

	err, response := Beneficiary.Bill(payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	fmt.Printf("hello, world i'm out here \n")
}
