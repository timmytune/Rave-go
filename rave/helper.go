//  helper functions

package rave

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
)

// Converts map[string]interface{} to JSON
func MapToJSON(mapData interface{}) ([]byte, error) {
	jsonBytes, err := json.Marshal(mapData)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

// Checks if all required parameters are present
func CheckRequiredParameters(params map[string]interface{}, keys []string) error {
	for _, key := range keys {
		if _, ok := params[key]; !ok {
			pc := make([]uintptr, 10)
			runtime.Callers(2, pc)
			f := runtime.FuncForPC(pc[0]).Name()
			details := strings.Split(f, ".")
			funcName := details[len(details)-1]
			return fmt.Errorf("%s is a required parameter for %s", key, funcName)
		}
	}
	return nil
}

// Makes a post request to rave api
func MakePostRequest(data interface{}, url string) (err error, response map[string]interface{}) {
	postData, err := MapToJSON(data)
	if err != nil {
		return err, nil
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postData))
	if err != nil {
		return err, nil
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return nil, result

}

func Request(typ string, data interface{}, url string, params map[string]string, query map[string]string, secretKey string) (response map[string]interface{}, err error) {
	postData, err := MapToJSON(data)
	if err != nil {
		return nil, err
	}

	if secretKey == "" {
		secretKey = rave.SecretKey
	}

	if len(params) > 0 {
		for k, v := range params {
			url = strings.ReplaceAll(url, k, v)
		}
	}

	if len(query) > 0 {
		var addToUrl string = "?"
		for k, v := range query {
			addToUrl += fmt.Sprintf("%s=%s&", k, v)
		}
		url += addToUrl
	}

	log.Print("URL: ", url)

	client := http.Client{}
	req, err := http.NewRequest(typ, url, bytes.NewBuffer(postData))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+secretKey)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func PostRequest(data interface{}, url string, params map[string]string, query map[string]string, secretKey string) (response map[string]interface{}, err error) {
	return Request("POST", data, url, params, query, secretKey)
}

func GetRequest(url string, params map[string]string, query map[string]string, secretKey string) (response map[string]interface{}, err error) {
	return Request("GET", nil, url, params, query, secretKey)
}

// Makes a get request to rave api
func MakeGetRequest(url string, params map[string]string) (error error, response map[string]interface{}) {
	var addToUrl string = "?"
	for k, v := range params {
		addToUrl += fmt.Sprintf("%s=%s&", k, v)
	}
	url += addToUrl
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return err, noresponse
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return nil, result

}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
func GenerateRef() string {
	len := 10
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return "MC-" + string(bytes)

}

// Checks that the transaction reference(TxRef) match
func VerifyTransactionReference(apiTransactionRef, funcTransactionRef interface{}) error {
	if apiTransactionRef != funcTransactionRef {
		return fmt.Errorf(
			"Transaction not verified because the transaction reference doesn't match: '%s' != '%s'",
			apiTransactionRef, funcTransactionRef,
		)
	}

	return nil
}

// The status should equal "success" for a succesful transaction
func VerifySuccessMessage(status string) error {
	if status != "success" {
		return fmt.Errorf("Transaction not verified because status is not equal to 'success'")
	}

	return nil
}

// The Charge response should equal "00" or "0"
func VerifyChargeResponse(chargeResponse string) error {
	if chargeResponse != "00" && chargeResponse != "0" {
		return fmt.Errorf("Transaction not verified because the charged response is not equal to '00' or '0'")
	}

	return nil
}

// The Currency code must match
func VerifyCurrencyCode(apiCurrencyCode, funcCurrencyCode interface{}) error {
	if apiCurrencyCode != funcCurrencyCode {
		return fmt.Errorf(
			"Transaction not verified because the currency code doesn't match: '%s' != '%s'",
			apiCurrencyCode, funcCurrencyCode,
		)
	}

	return nil
}

// The Charged Amount must be greater than or equal to the paid amount
func VerifyChargedAmount(apiChargedAmount, funcChargedAmount float64) error {
	if funcChargedAmount < apiChargedAmount {
		return fmt.Errorf("Transaction not verified, incorrect amount: charged amount should be greater or equal amount to be paid")
	}

	return nil
}
