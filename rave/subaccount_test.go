package rave

import (
	"testing"
)

var r = Rave{
	false,
	"FLWPUBK-xxxxxxxxxxxxxxxxxxxxx-X",
	"FLWSECK-xxxxxxxxxxxxxxxxxxxxx-X",
}

func TestCreateSubaccount(t *testing.T) {
	var tests = []CreateSubaccountData{
		{
			AccountBank:           "044",
			AccountNumber:         "0690000035",
			BusinessName:          "AJB",
			BusinessEmail:         "jk@services.com",
			BusinessContact:       "gigi alade",
			BusinessContactMobile: "09012345678",
			BusinessMobile:        "09087930123",
			SplitType:             "flat",
			SplitValue:            "100",
		},
	}

	for _, test := range tests {
		error, response := Subaccount{
			r,
		}.CreateSubaccount(test)
		if error != nil {
			t.Fatalf("Create Subaccount failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Create Subaccount status: %v, Details: %v", response["status"], response)
		}
	}
}

func TestListSubaccount(t *testing.T) {
	var tests = []ListSubaccountData{
		{
			AccountNumber: "0690000035",
		},
	}

	for _, test := range tests {
		error, response := Subaccount{
			r,
		}.ListSubaccount(test)
		if error != nil {
			t.Fatalf("List Subaccount failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("List Subaccount status: %v, Details: %v", response["status"], response)
		}
	}
}

func TestFetchSubaccount(t *testing.T) {
	ids := []string{"RS_F2CF2DC6FA3A67BD4E74309F3FE3594C"}

	for _, id := range ids {
		error, response := Subaccount{
			r,
		}.FetchSubaccount(id)
		if error != nil {
			t.Fatalf("Fetch Subaccount failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Fetch Subccount status: %v, Details: %v", response["status"], response)
		}
	}
}

func TestDeleteSubaccount(t *testing.T) {
	ids := []string{"RS_F2CF2DC6FA3A67BD4E74309F3FE3594C"}

	for _, id := range ids {
		error, response := Subaccount{
			r,
		}.DeleteSubaccount(id)
		if error != nil {
			t.Fatalf("Delete Subaccount failed with error %v", error)
		}
		if response["status"] != "success" {
			t.Fatalf("Delete Subccount status: %v, Details: %v", response["status"], response)
		}
	}
}
