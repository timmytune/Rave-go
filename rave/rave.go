package rave

import (
	"bytes"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"log"
	"os"
	"strings"
)

const (
	sandboxUrl string = "https://api.flutterwave.com/"
	liveUrl    string = "https://api.flutterwave.com/"
)

// Rave base type
type Rave struct {
	Live      bool
	PublicKey string
	SecretKey string
}

var rave Rave

func InitFlutterwave(isLive bool, secretKey string, publicKey string) {
	if secretKey == "" || publicKey == "" {
		log.Fatal("You need to set the your publickey and secretKey")
	}
	rave = Rave{Live: isLive, PublicKey: publicKey, SecretKey: secretKey}
}

var Endpoints = map[string]map[string]string{
	"card": {
		"charge":          "flwv3-pug/getpaidx/api/charge",
		"validate":        "flwv3-pug/getpaidx/api/validatecharge",
		"verify":          "flwv3-pug/getpaidx/api/v2/verify",
		"chargeSavedCard": "flwv3-pug/getpaidx/api/tokenized/charge",
	},
	"preauth": {
		"charge":       "flwv3-pug/getpaidx/api/tokenized/preauth_charge",
		"capture":      "flwv3-pug/getpaidx/api/capture",
		"refundorvoid": "flwv3-pug/getpaidx/api/refundorvoid",
	},
	"account": {
		"charge":   "flwv3-pug/getpaidx/api/charge",
		"validate": "flwv3-pug/getpaidx/api/validate",
		"verify":   "flwv3-pug/getpaidx/api/v2/verify",
	},
	"payment_plan": {
		"create": "v2/gpx/paymentplans/create",
		"fetch":  "v2/gpx/paymentplans/query",
		"list":   "v2/gpx/paymentplans/query",
		"cancel": "v2/gpx/paymentplans/",
		"edit":   "v2/gpx/paymentplans/",
	},
	"subscriptions": {
		"fetch":    "v2/gpx/subscriptions/query",
		"list":     "v2/gpx/subscriptions/query",
		"cancel":   "v2/gpx/subscriptions/",
		"activate": "v2/gpx/subscriptions/",
	},
	"subaccount": {
		"create": "v2/gpx/subaccounts/create",
		"list":   "v2/gpx/subaccounts/",
		"fetch":  "v2/gpx/subaccounts/get",
		"delete": "v2/gpx/subaccounts/delete",
	},
	"transfer": {
		"initiate":            "v2/gpx/transfers/create",
		"bulk":                "v2/gpx/transfers/create_bulk",
		"fetch":               "v2/gpx/transfers",
		"fee":                 "v2/gpx/transfers/fee",
		"balance":             "v2/gpx/balance",
		"accountVerification": "flwv3-pug/getpaidx/api/resolve_account",
	},
	"verify": {
		"verify": "flwv3-pug/getpaidx/api/v2/verify",
	},
	"refund": {
		"refund": "gpx/merchant/transactions/refund",
	},
	"mobilemoney": {
		"charge": "flwv3-pug/getpaidx/api/charge",
	},
	"virtualcard": {
		"create":    "v2/services/virtualcards/new",
		"list":      "v2/services/virtualcards/search",
		"get":       "v2/services/virtualcards/get",
		"terminate": "v2/services/virtualcards/",
		"fund":      "v2/services/virtualcards/fund",
		"fetch":     "v2/services/virtualcards/transactions",
		"withdraw":  "v2/services/virtualcards/withdraw",
		"freeze":    "v2/services/virtualcards/",
	},

	"bvn": {
		"bvnverification": "v2/kyc/bvn/",
	},

	"virtualaccount": {
		"virtualaccountnumber": "v2/banktransfers/accountnumbers",
	},

	"settlement": {
		"list":  "v2/merchant/settlements",
		"fetch": "v2/merchant/settlements/",
	},

	"verifytransaction": {
		"verify": "flwv3-pug/getpaidx/api/v2/verify",
	},

	"ebills": {
		"createorder": "flwv3-pug/getpaidx/api/ebills/generateorder/",
		"updateorder": "flwv3-pug/getpaidx/api/ebills/update/",
	},

	"Banktransfer": {
		"charge": "flwv3-pug/getpaidx/api/charge",
	},

	"Beneficiaries": {
		"list":   "v2/gpx/transfers/beneficiaries",
		"fetch":  "v2/gpx/transfers/beneficiaries",
		"create": "v2/gpx/transfers/beneficiaries/create",
		"delete": "v2/gpx/transfers/beneficiaries/delete",
	},
	"Billspayments": {
		"flybuy": "v2/services/confluence",
	},
	"flutterwaveOTP": {
		"otp": "v2/services/confluence",
	},
	"transaction": {
		"get.id":   "https://api.flutterwave.com/v3/transactions/:id/verify",
		"get.many": "https://api.flutterwave.com/v3/transactions",
	},
}

// gets the correct url for live and test mode
func (r Rave) GetBaseURL() string {
	if r.Live {
		return liveUrl
	}

	return sandboxUrl
}

func (Rave) GetEndpoint(raveType string, action string) string {
	return Endpoints[raveType][action]
}

// gets the public key from the environment variable if set or from the Rave object
func (r Rave) GetPublicKey() string {
	if r.PublicKey != "" {
		return r.PublicKey
	}
	pubKey, ok := os.LookupEnv("RAVE_PUBKEY")
	if !ok {
		log.Fatal("You need to set the your public key as an environment variable \"RAVE_PUBKEY\" or as a field in the Rave struct")
	}
	r.PublicKey = pubKey
	return pubKey
}

// gets the secret key
func (r Rave) GetSecretKey() string {
	if r.SecretKey != "" {
		return r.SecretKey
	}
	secKey, ok := os.LookupEnv("RAVE_SECKEY")
	if !ok {
		log.Fatal("You need to set the your secret key as an environment variable \"RAVE_SECKEY\" or as a field in the Rave struct")
	}
	r.SecretKey = secKey
	return secKey
}

func (r Rave) getKey(seckey string) string {
	keymd5 := md5.Sum([]byte(seckey))
	keymd5Last12 := keymd5[len(keymd5)-6:] // -6 because it's a hex byte array not a string
	seckeyAdjusted := strings.Replace(seckey, "FLWSECK-", "", 1)
	seckeyAdjustedFirst12 := seckeyAdjusted[:12]

	return seckeyAdjustedFirst12 + hex.EncodeToString(keymd5Last12[:])
}

func (r Rave) Encrypt(payload string) string {
	seckey := r.GetSecretKey()
	encryptedSecKey := r.getKey(seckey)

	return r.encrypt3Des(encryptedSecKey, payload)

}

// Encrypt3Des : Encrypts the data using 3Des encryption
func (r Rave) encrypt3Des(key string, payload string) string {
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	bs := block.BlockSize() // block size is 8 by default
	payloadBytes := pkcs5Padding([]byte(payload), bs)

	if len(payloadBytes)%bs != 0 {
		panic("Need a multiple of the blocksize")
	}
	encrypted := make([]byte, len(payloadBytes))
	dst := encrypted

	for len(payloadBytes) > 0 {
		block.Encrypt(dst, payloadBytes[:bs])
		payloadBytes = payloadBytes[bs:]
		dst = dst[bs:]
	}

	return base64.StdEncoding.EncodeToString(encrypted)
}

// pkcs5Padding : Implements PKCS5 padding
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
