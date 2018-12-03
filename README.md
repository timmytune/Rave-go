# Rave-go

## Introduction
This is a Go wrapper around the [API](https://flutterwavedevelopers.readme.io/v2.0/reference) for [Rave by Flutterwave](https://rave.flutterwave.com).
#### Payment types implemented:
* Card Payments
* Bank Account Payments
* Preauth
* Refund
* Subaccount
* Transfer
* Payment Plan
* Subscription
* USSD Payments
## Installation
To install, run

``` go get github.com/anjolabassey/rave```

Note: This is currently under active development
## Import Package
The base class for this package is 'Rave'. To use this class, add:

```
 import (
 	"github.com/anjolabassey/Rave-go/rave"
 )
 ```

## Initialization

#### To instantiate in sandbox:
To use Rave, instantiate Rave with your public key. We recommend that you store your secret key in an environment variable named, ```RAVE_SECKEY```. However, you can also pass it in here alongside your public key. Instantiating Rave is therefore as simple as:


```
var r = rave.Rave{
	false,
	"FLWPUBK-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-X",
	"FLWSECK-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-X",
}
```

**Note: If you store your secret key as an environment variable, just pass an empty string "" for the secret field as shown below**

```
var r = rave.Rave{
	false,
	"FLWPUBK-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-X",
	"",
}
```

#### To instantiate in production:
To initialize in production, simply set the ```production``` flag to ```true```.

```
var r = rave.Rave{
	true,
	"FLWPUBK-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-X",
	"FLWSECK-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-X",
}
```

# Rave Methods
This is the documentation for all of the components of Rave-go

## ```rave.Card{}```
This is used to facilitate card charges via rave. ```rave.Card{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```
    var card = rave.Card{
    	r,
    }
```
**Methods Included:**

* ```.ChargeCard```

* ```.ValidateCard```

* ```.VerifyCard```

* ```.TokenizedCharge```

* ```.ChargePreauth```

* ```.CapturePreauth```

### ```.ChargeCard(data CardChargeData) (error error, response map[string]interface{})```
This is called to charge a card. The payload should be of type ```rave.CardChargeData```. See below for  ```rave.CardChargeData``` definition

```
type CardChargeData struct {
	Cardno               string         `json:"cardno"`
	Cvv                  string         `json:"cvv"`
	Expirymonth          string         `json:"expirymonth"`
	Expiryyear           string         `json:"expiryyear"`
	Pin                  string         `json:"pin"`
	Amount               float64        `json:"amount"`
	Currency             string         `json:"currency"`
	Country              string         `json:"country"`
	CustomerPhone        string         `json:"customer_phone"`
	Firstname            string         `json:"firstname"`
	Lastname             string         `json:"lastname"`
	Email                string         `json:"email"`
	Ip                   string         `json:"IP"`
	Txref		         string	        `json:"txRef"`
	RedirectUrl          string         `json:"redirect_url"`
	Subaccounts          types.Slice    `json:"subaccounts"`
	DeviceFingerprint    string         `json:"device_fingerprint"`
	Meta                 types.Slice    `json:"meta"`
	SuggestedAuth        string         `json:"suggested_auth"`
	BillingZip           string         `json:"billingzip"`
	BillingCity          string         `json:"billingcity"`
	BillingAddress       string         `json:"billingaddress"`
	BillingState         string         `json:"billingstate"`
	BillingCountry       string         `json:"billingcountry"`
	Chargetype		     string	        `json:"charge_type"`

}
```
A sample initiate call is:

```
    payload := rave.CardChargeData{

        Amount:100,
		Txref:"MC-11001993",
		Email:"test@test.com",
		CustomerPhone:"08123456789",
		Currency:"NGN",
		Cardno:"5399838383838381",
		Cvv:"470",
		Expirymonth:"10",
		Expiryyear:"22",
		Pin: "3310",
    }

    err, response := card.ChargeCard(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```

```

### ```.ValidateCard(data CardValidateData) (error error, response map[string]interface{})```
This is called to validate a card charge. The payload should be of type ```rave.CardValidateData```. See below for  ```rave.CardValidateData``` definition
```
type CardValidateData struct {
	Reference	   string	      `json:"transaction_reference"`
	Otp		       string	      `json:"otp"`
	PublicKey      string         `json:"PBFPubKey"``
}
The Reference is the `flwRef` gotten from the response of the ChargeCard function. See an example below
ref := response["data"].(map[string]interface{})["flwRef"].(string)
```
A sample initiate call is:

```
    payload := rave.CardValidateData{
        Otp:"12345",
		Reference: ref,
    }

    err, response := card.ValidateCard(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```

```

### ```.VerifyCard(data CardVerifyData) (error error, response map[string]interface{})```
This is called to validate a card charge. The payload should be of type ```rave.CardVerifyData```. See below for  ```rave.CardVerifyData``` definition
```
type CardVerifyData struct {
	Reference	   string	      `json:"transaction_reference"`
	Otp		       string	      `json:"otp"`
	PublicKey      string         `json:"PBFPubKey"``
}
The Reference is the `txRef` which can also be gotten from the response of the ChargeCard function. See example below
txref := response["data"].(map[string]interface{})["txRef"].(string)
```
A sample initiate call is:

```
    payload := rave.CardVerifyData{
        Otp:"12345",
		Reference: txref,
    }

    err, response := card.VerifyCard(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```

```
### ```.TokenizedCharge(data SaveCardChargeData) (error error, response map[string]interface{})```
This is called to charge a saved card using a token(which can be gotten in the [verify payment response](https://developer.flutterwave.com/v2.0/reference#save-a-card)). The payload should be of type ```rave.TokenizedChargeData```. See below for  ```rave.TokenizedChargeData``` definition
```
type TokenizedChargeData struct {
    SecretKey      string         `json:"SECKEY"`
	Currency       string         `json:"currency"`
	Token          string         `json:"token"`
	Country        string         `json:"country"`
	Amount	       float64	      `json:"amount"`
	Email          string         `json:"email"`
	Firstname      string         `json:"firstname"`
	Lastname       string         `json:"lastname"`
	Ip             string         `json:"IP"`
	Txref		   string	      `json:"txRef"`
}

```
A sample initiate call is:

```
    payload := rave.TokenizedChargeData{
        Token: "flw-t1nf-2f00ba4c24b27cbb39e7907c6b72d413-m03k",
		Currency:"NGN",
	 	Country:"NG",
	 	Amount:100,
	 	Email:"test@test.com",
	 	Txref:"MC-0123456789",
    }

    err, response := card.TokenizedCharge(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```

```

## ```rave.Refund{}```
This allows you initiate refunds for Successful transactions via rave. ```rave.Refund{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```
    var transfer = rave.Refund{
    	r,
    }
```

**Methods Included:**

* ```.RefundTransaction```

### ```.RefundTransaction(data RefundData) (error error, response map[string]interface{})```
This is called to initiate refunds for Successful transaction. The payload should be of type ```rave.RefundData```. See below for  ```rave.RefundData``` definition

```
type RefundData struct {
	Ref		       string	      `json:"ref"`
	Amount         int            `json:"amount"`
	SecretKey      string         `json:"seckey"`
}
```

A sample refund call is:

```
    payload := rave.RefundData{
        Ref: "FLW-MOCK-476a260e67df43988a2ffeddf8e02cc2",
		Amount: 100,
    }

    err, response := refund.RefundTransaction(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```

```

## ```rave.Preauth{}```
This is used to facilitate preauthorised card transactions via rave. ```rave.Preauth{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property. This inherits the Card class so any task you can do on Card, you can do with preauth.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```
    var card = rave.Preauth{
    	r,
    }
```
**Methods Included:**

* ```.ChargePreauth```

* ```.VerifyPreauth```

* ```.PreauthTokenizedCharge```

* ```.CapturePreauth```

* ```.RefundOrVoidPreauth```

### ```.CapturePreauth(data PreauthCaptureData) (error error, response map[string]interface{})```
This is called to preauthorize a card. The payload should be of type ```rave.PreauthCaptureData```. 

A sample initiate call is:

```
    payload := rave.PreauthCaptureData{
        Amount:100,
        Flwref: ref

    }

    err, response := card.CapturePreauth(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```

```

### ```.ChargePreauth(data CardChargeData) (error error, response map[string]interface{})```
This is called to preauthorize a card. The payload should be of type ```rave.CardChargeData```. 

A sample initiate call is:

```
    payload := rave.CardChargeData{
        Amount:100,
		Txref:"MC-11001993",
		Email:"test@test.com",
		CustomerPhone:"08123456789",
		Currency:"NGN",
		Cardno:"5399838383838381",
		Cvv:"470",
		Expirymonth:"10",
		Expiryyear:"22",
		Pin: "3310",
    }

    err, response := card.ChargePreauth(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```
#### Sample Response

```

```

## ```rave.Transfer{}```
This is used to facilitate transfers via rave. ```rave.Transfer{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```
    var transfer = rave.Transfer{
    	r,
    }
```

**Methods Included:**

* ```.InitiateSingleTransfer```

* ```.InitiateBulkTransfer```

* ```.FetchTransfer```

* ```.FetchAllTransfers```

* ```.GetTransferFee```

* ```.GetRaveBalance```

* ```.GetBulkTransferStatus```

### ```.InitiateSingleTransfer(payload SinglePaymentData) (error error, response map[string]interface{})```
This is called to initiate a sole transfer. The payload should be of type ```rave.SinglePaymentData```. See below for  ```rave.SinglePaymentData``` definition

```
type SinglePaymentData struct {
	SecKey          string      `json:"seckey"`
	AccountBank     string      `json:"account_bank"`
	AccountNumber   string      `json:"account_number"`
	Amount          int         `json:"amount"`
	Narration       string      `json:"narration"`
	Currency        string      `json:"currency"`
	Reference       string      `json:"reference"`
	Meta            types.Slice `json:"meta"`
	BeneficiaryName string      `json:"beneficiary_name"`
}
```

A sample initiate call is:

```
    payload := rave.SinglePaymentData{
        AccountBank: "044",
        AccountNumber: "0690000044",
        Amount:        500,
        SecKey:        r.GetSecretKey(),
        Narration:     "Test Transfer",
        Currency:      "NGN",
        Reference:     time.Now().String(),
    }

    err, response := transfer.InitiateSingleTransfer(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```
map[status:success message:TRANSFER-CREATED data:map[id:3603 account_number:0690000044 bank_code:044 date_created:2018-11-27T12:15:37.000Z amount:500 currency:NGN meta:map[] is_approved:1 bank_name:ACCESS BANK NIGERIA fullname:Mercedes Daniel status:NEW reference:2018-11-27 13:15:36.5762772 +0100 WAT m=+0.013132801 narration:Test Transfer requires_approval:0 fee:45 complete_message:]]
```

### ```.InitiateBulkTransfer(payload BulkPaymentData) (error error, response map[string]interface{})```

This is called to initiate a bulk transfer. The payload should be of type ```rave.BulkPaymentData```. See below for  ```rave.BulkPaymentData``` definition

```
type BulkPaymentData struct {
    SecKey   string              `json:"seckey"`
    Title    string              `json:"title"`
    BulkData []map[string]string `json:"bulk_data"`
}
```

A sample initiate call is:

```
    payloads := rave.BulkPaymentData{
        SecKey: "FLWSECK-0b1d6669cf375a6208db541a1d59adbb-X",
        Title:  "May Staff Salary",
        BulkData: []map[string]string{
            {
                "Bank":           "044",
                "Account Number": "0690000032",
                "Amount":         "500",
                "Currency":       "NGN",
                "Narration":      "Bulk transfer 1",
                "reference":      time.Now().String(),
            },
            {
                "Bank":           "044",
                "Account Number": "0690000034",
                "Amount":         "500",
                "Currency":       "NGN",
                "Narration":      "Bulk transfer 2",
                "reference":      time.Now().String(),
            },
        },
    }


    err, response := transfer.InitiateBulkTransfer(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```
map[status:success message:BULK-TRANSFER-CREATED data:map[id:683 date_created:2018-11-27T12:57:59.000Z approver:N/A]]
```


### ```.FetchTransfer(reference string) (error error, response map[string]interface{})```

This allows you retrieve a single transfer. The reference should be of type ```string```You may or may not pass in a transfer ```reference```. If you do not pass in a reference, all transfers that have been processed will be returned.

A sample fetch call is:

```
reference := "kkkkkkkkkkkkk"
err, response := transfer.FetchTransfer(reference)
if err != nil {
 panic(err)
}
fmt.Println(response)
```

#### Sample Response

This call returns a dictionary. A sample response is:

 ```
 map[status:success message:QUERIED-TRANSFERS data:map[page_info:map[total:1 current_page:1 total_pages:1] transfers:[map[requires_approval:0 debit_currency:<nil> reference:kkkkkkkkkkkkk complete_message:DISBURSE FAILED: undefined is_approved:1 id:3563 account_number:0690000044 bank_code:044 currency:NGN fee:45 meta:map[] approver:<nil> bank_name:ACCESS BANK NIGERIA narration:New transfer fullname:Mercedes Daniel date_created:2018-11-26T21:38:31.000Z amount:500 status:FAILED]]]]
 ```

 ### ```.FetchTransfer(reference string) (error error, response map[string]interface{})```

 This allows you retrieve a single transfer. The reference should be of type ```string```You may or may not pass in a transfer ```reference```. If you do not pass in a reference, all transfers that have been processed will be returned.

 A sample fetch call is:

 ```
 reference := "kkkkkkkkkkkkk"
 err, response := transfer.FetchTransfer(reference)
 if err != nil {
  panic(err)
 }
 fmt.Println(response)
 ```

 #### Sample Response

 This call returns a dictionary. A sample response is:

  ```
  map[status:success message:QUERIED-TRANSFERS data:map[page_info:map[total:1 current_page:1 total_pages:1] transfers:[map[requires_approval:0 debit_currency:<nil> reference:kkkkkkkkkkkkk complete_message:DISBURSE FAILED: undefined is_approved:1 id:3563 account_number:0690000044 bank_code:044 currency:NGN fee:45 meta:map[] approver:<nil> bank_name:ACCESS BANK NIGERIA narration:New transfer fullname:Mercedes Daniel date_created:2018-11-26T21:38:31.000Z amount:500 status:FAILED]]]]
  ```

### ```.FetchAllTransfers() (error error, response map[string]interface{})```

This allows you retrieve all transfers.

A sample fetchall call is:

```
err, response := transfer.FetchAllTransfers()
if err != nil {
 panic(err)
}
fmt.Println(response)
```

#### Sample Response

This call returns a dictionary. A sample response is:

 ```
 map[status:success message:QUERIED-TRANSFERS data:map[page_info:map[total:1 current_page:1 total_pages:1] transfers:[map[requires_approval:0 debit_currency:<nil> reference:kkkkkkkkkkkkk complete_message:DISBURSE FAILED: undefined is_approved:1 id:3563 account_number:0690000044 bank_code:044 currency:NGN fee:45 meta:map[] approver:<nil> bank_name:ACCESS BANK NIGERIA narration:New transfer fullname:Mercedes Daniel date_created:2018-11-26T21:38:31.000Z amount:500 status:FAILED]]]]
 ```


### ```.GetTransferFee(currency string) (error error, response map[string]interface{})```

This allows you get transfer rates for all Rave supported currencies. You may or may not pass in a ```currency```. If you do not pass in a ```currency```, all Rave supported currencies transfer rates will be returned.

A sample getFee call is:

```
currencies := "NGN"
error, response := transfer.GetTransferFee(currency)
if err != nil {
 panic(err)
}
fmt.Println(response)

```

#### Sample Response

This call returns a dictionary. A sample response is:

 ```
map[message:TRANSFER-FEES data:[map[AccountId:1 id:1 fee_type:value currency:NGN fee:45 createdAt:<nil> updatedAt:<nil> deletedAt:<nil>]] status:success]

 ```

### ```.GetRaveBalance(currency string) (error error, response map[string]interface{})```

This allows you get your balance in a specified currency. You may or may not pass in a ```currency```. If you do not pass in a ```currency```, your balance will be returned in the currency specified in yiur rave account

A sample balance call is:

```
currencies := "NGN"
error, response := transfer.GetRaveBalance(currency)
if err != nil {
 panic(err)
}
fmt.Println(response)
```

#### Returns

This call returns a dictionary. A sample response is:

 ```
map[status:success message:WALLET-BALANCE data:map[LedgerBalance:0 Id:32509 ShortName:NGN WalletNumber:4446000147772 AvailableBalance:0]]
 ```

### ```.GetBulkTransferStatus(batch_id string) (error error, response map[string]interface{})```

This allows you get your status of a queued bulk transfer You may or may not pass in a ```batch_id```. If you do not pass in a ```batch_id```, all queued bulk transfers will be returned

A sample bulk transfer status call is:

```
batchIDs := [2]string{"634", "635"}

error, response := transfer.GetBulkTransferStatus(batchID)
if err != nil {
 panic(err)
}
fmt.Println(response)
```

#### Returns

This call returns a dictionary. A sample response is:

 ```
map[message:QUERIED-TRANSFERS data:map[page_info:map[total_pages:1 total:2 current_page:1] transfers:[map[bank_name:ACCESS BANK NIGERIA account_number:0690000032 fullname:Pastor Bright amount:10 reference:<nil> narration:Bulk transfer 1 approver:<nil> id:3542 bank_code:044 requires_approval:0 is_approved:1 date_created:2018-11-26T14:21:44.000Z currency:NGN debit_currency:<nil> fee:45 status:FAILED meta:<nil> complete_message:DISBURSE FAILED: Invalid transfer amount. Minimum is 100] map[bank_code:044 fee:45 meta:<nil> account_number:0690000034 currency:NGN debit_currency:<nil> amount:10 complete_message:DISBURSE FAILED: Invalid transfer amount. Minimum is 100 fullname:Ade Bond date_created:2018-11-26T14:21:44.000Z reference:<nil> narration:Bulk transfer 2 approver:<nil> is_approved:1 id:3543 status:FAILED requires_approval:0 bank_name:ACCESS BANK NIGERIA]]] status:success]
 ```


### ```.ResolveAccount(account_data AccountResolveData) (error error, response map[string]interface{})```
This allows you verify an account to transfer to. ```account_data``` should be of type ```rave.AccountResolveData```. See below for  ```rave.AccountResolveData``` definition

```
type AccountResolveData struct {
	PublicKey        string `json:"PBFPubKey"`
	RecipientAccount string `json:"recipientaccount"`
	DestBankCode     string `json:"destbankcode"`
	Currency         string `json:"currency"`
	Country          string `json:"country"`
}
```

A sample initiate call is:

```
    payload := rave.AccountResolveData{
        RecipientAccount: "0690000034",
        DestBankCode:     "044",
        PublicKey:        r.GetPublicKey(),
    }

    err, response := transfer.ResolveAccount(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```
map[message:ACCOUNT RESOLVED data:map[data:map[responsecode:00 accountnumber:0690000034 accountname:Ade Bond responsemessage:Approved Or Completed Successfully phonenumber:<nil> uniquereference:FLWT001034195 internalreference:<nil>] status:success] status:success]
```

## ```rave.Subaccount{}```
This is used to facilitate Subaccount operations via rave. ```rave.Subaccount{}``` is of type ```struct``` and requires  ```rave.Rave``` as its only property.

Hence, in order to use it, you need to pass in an instance of ```rave.Rave``` . A sample is shown below

```
    var transfer = rave.Subaccount{
    	r,
    }
```

**Methods Included:**

* ```.CreateSubaccount```

* ```.ListSubaccount```

* ```.FetchSubaccount```

* ```.DeleteSubaccount```

### ```.CreateSubaccount(data CreateSubaccountData) (error error, response map[string]interface{})```
This is called to create a subaccount. The payload should be of type ```rave.CreateSubaccountData```. See below for  ```rave.CreateSubaccountData``` definition

```
type CreateSubaccountData struct {
	AccountBank                   string         `json:"account_bank"`
	AccountNumber                 string         `json:"account_number"`
	BusinessName                  string         `json:"business_name"`
	BusinessEmail                 string         `json:"business_email"`
	BusinessContact               string         `json:"business_contact"`
	BusinessMobile                string         `json:"business_mobile"`
	BusinessContactMobile         string         `json:"business_contact_mobile"`
	Seckey                        string         `json:"seckey"`
	Meta                          types.Slice    `json:"meta"`
	SplitType                     string         `json:"split_type"`
	SplitValue                    string         `json:"split_value"`
}
```

A sample create call is:

```
    payload := rave.CreateSubaccountData{
        AccountBank: "044",
	 	AccountNumber: "0690000035",
	 	BusinessName: "Test",
	 	BusinessEmail: "test@test.com",
	 	BusinessContact: "Seun Alade",
	 	BusinessContactMobile: "09012345678",
	 	BusinessMobile: "09087930123",
	 	SplitType: "flat",
	 	SplitValue: "100",
    }

    err, response := subaccount.CreateSubaccount(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```

```
### ```.ListSubaccount(data ListSubaccountData) (error error, response map[string]interface{})```
This is called to list all or specific subaccounts. The payload should be of type ```rave.ListSubaccountData```. See below for  ```rave.ListSubaccountData``` definition

```
type ListSubaccountData struct {
	AccountBank               string         `json:"account_bank"`
	AccountNumber             string         `json:"account_number"`
	BankName                  string         `json:"bank_name"`
	Seckey                    string         `json:"seckey"`
}
```

A sample list call is:

```
    payload := rave.ListSubaccountData{
        AccountNumber: "0690000035",
    }

    err, response := subaccount.ListSubaccount(payload)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```

```
### ```.FetchSubaccount(id string) (error error, response map[string]interface{})```
This allows you fetch a single subaccount using the subaccount ID. The ID should be of type ```string``` and it is required.

A sample fetch call is:


```
    id = "example_id"

    err, response := subaccount.FetchSubaccount(id)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```

```
### ```.DeleteSubaccount(id string) (error error, response map[string]interface{})```
This allows you to delete a subaccount by the subaccount ID. The ID should be of type ```string``` and it is required.

A sample delete call is:


```
    id = "example_id"

    err, response := subaccount.DeleteSubaccount(id)
    if err != nil {
        panic(err)
    }
    fmt.Println(response)
```

#### Sample Response

```

```
