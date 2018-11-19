package main

import (
	"Rave-go/rave"
	"Rave-go/rave/refund"
	"fmt"
)

func main() {
	r := rave.Rave{
		false,
		"FLWPUBK-f54d8d24292e377a71620bd82a8bb17c-X",
		"FLWSECK-a18ca169cb007a93db4479aff683a387-X",
	}
	c := refund.Refund{
		r,
	}
	 a := refund.RefundData {
		Ref: "FLW-MOCK-476a260e67df43988a2ffeddf8e02cc2",
		Amount: 1, 
		
	 }
	err, response := c.RefundTransaction(a)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(response)
	}
