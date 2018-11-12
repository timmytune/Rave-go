package rave

import (
	"os"
	"log"
)

// Rave base type
type Rave struct {
	Live bool
	liveUrl string
	testUrl string
	PublicKey string
	SecretKey string
}

// gets the correct url for live and test mode
func (r Rave) getBaseURL() string {
	if r.Live {
		return r.liveUrl
	}

	return r.testUrl
}

// gets the public key
func (r Rave) getPublicKey() string {
	pubKey, ok := os.LookupEnv("RAVE_PUBKEY")
	if !ok {
		log.Fatal("You need to set the \"RAVE_PUBKEY\" environment variable")
	}
	return pubKey
}

// gets the secret key
func (r Rave) getSecretKey() string {
	secKey, ok := os.LookupEnv("RAVE_SECKEY")
	if !ok {
		log.Fatal("You need to set the \"RAVE_SECKEY\" environment variable")
	}
	return secKey
}

// constructor for Rave staruct
func NewRave() Rave {
	Rave := Rave{}
	Rave.testUrl = "https://ravesandboxapi.flutterwave.com"
	Rave.liveUrl = "https://api.ravepay.co"
	Rave.Live = false

	return Rave
}