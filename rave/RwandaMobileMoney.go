package rave

// "fmt"
// "go/types"

//Sub Interface

type RwandaMobilemoneyCharge interface {
	RwandaMobilemoneyCharge(data RwandaMobileChargeData) (error error, response map[string]interface{})
}

//Main Interface

type Rwandamobilemoney interface {
	RwandaMobilemoneyCharge
}
type RwandaMobileChargeData struct {
	Pubkey            string `json:"PBFPubKey"`
	Currency          string `json:"currency"`
	PaymentType       string `json:"payment_type"`
	Country           string `json:"country"`
	Amount            string `json:"amount"`
	Email             string `json:"email"`
	Phonenumber       string `json:"phonenumber"`
	Network           string `json:"network"`
	Firstname         string `json:"firstname"`
	Lastname          string `json:"lastname"`
	Voucher           string `json:"voucher"`
	IP                string `json:"IP"`
	Txref             string `json:"txRef"`
	OrderRef          string `json:"orderRef"`
	IsMobile          int    `json:"is_mobile_money_gh"`
	RedirectURL       string `json:"redirect_url"`
	DeviceFingerprint string `json:"device_fingerprint"`
}

type RwandaMobilemoney struct {
	Rave
}

func (m RwandaMobilemoney) SetupCharge(data RwandaMobileChargeData) map[string]interface{} {
	chargeJSON, _ := MapToJSON(data)
	encryptedChargeData := m.Encrypt(string(chargeJSON[:]))
	queryParam := map[string]interface{}{
		"PBFPubKey": m.GetPublicKey(),
		"client":    encryptedChargeData,
		"alg":       "3DES-24",
	}
	return queryParam
}

func (m RwandaMobilemoney) RwandaMobilemoneyCharge(data RwandaMobileChargeData) (error error, response map[string]interface{}) {

	var url string
	url = m.GetBaseURL() + m.GetEndpoint("mobilemoney", "charge")
	if data.Txref == "" {
		data.Txref = GenerateRef()
	}
	postData := m.SetupCharge(data)

	err, response := MakePostRequest(postData, url)
	if err != nil {
		return err, noresponse
	}

	return nil, response

}
