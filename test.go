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
		"FLWPUBK_TEST-",
		"FLWSECK_TEST-9e",
	}
	var createotp = rave.FlutterwaveOTP{
		r,
	}

	CustomerInfoData:=rave.CustomerInfoData {
		Firstname: "Olufemi",
		Email:"olufemiobafunmiso@gmail.com",
		Mobile:"2347065489493",

	}



	servicepayloaddata := rave.ServicepayData{
		LenghtOfOTP:6,
		SendOTPasCustomer:true,
		CustomerInfo:CustomerInfoData,
		Medium: []string{"whatsapp"},
	}
	payload := rave.OTPData{
		Seckey:         "FLWSECK",
		Service:        "fly_otp_create",
		SenderAsIs:        1,
		ServiceMethod:  "post",
		ServiceVersion: "v1",
		ServiceChannel: "rave",
		ServicePayload: servicepayloaddata,
		OtpExpiresInMins: 6,
		SenderBusinessname: "O",
		SenderSameOTP:true,

	}

	fmt.Println(payload)


	err, response := createotp.Otp(payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
	fmt.Printf("hello, world i'm out here \n")
}

