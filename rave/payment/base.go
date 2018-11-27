package payment

import (
	// "Rave-go/rave/helper"
	"fmt"
	"math/rand"
	// "time"
	
)

// func SetupCharge(data) () {
// 	chargeJSON := helper.MapToJSON(data)
// 	encryptedChargeData := c.Encrypt(string(chargeJSON[:]))
// 	queryParam := map[string]interface{}{
//         "PBFPubKey": c.GetPublicKey(),
//         "client": encryptedChargeData,
//         "alg": "3DES-24",
//     }
// 	return queryParam
// }
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
		return fmt.Errorf("Transaction not verified, charged amount should be greater or equal amount to be paid")
	}

	return nil
}