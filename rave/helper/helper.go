//  helper functions

package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	// "io/ioutil"
	"runtime"
	// "github.com/antonholmquist/jason"
)

var noresponse = map[string]interface{}{
	"": "",
}

// Converts map[string]interface{} to JSON
func mapToJSON(mapData interface{}) []byte {
	jsonBytes, err := json.Marshal(mapData)
	if err != nil {
		panic(err)
	}

	return jsonBytes
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
			return fmt.Errorf("%s is a required parameter for %s\n", key, funcName)
		}
	}

	return nil
}

// Makes a post request to rave api
func MakePostRequest(data interface{}, url string) (error error, response map[string]interface{}) {
	postData := mapToJSON(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(postData))
	if err != nil {
		return err, noresponse
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	return nil, result

}

func MakeGetRequest(url string, params map[string]string) (error error, response map[string]interface{}) {
	var addToUrl string = "?"
	for k, v := range params {
		addToUrl += fmt.Sprintf("%s=%s&", k, v)
	}
	fmt.Println(addToUrl)
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

// // makes a post request to rave api
// func MakePostRequest(data map[string]interface{}, URL string) ([]byte) {
//     postData, err := json.Marshal(data)
//     if err != nil {
//         fmt.Println(err)
//     }

// 	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(postData))

//     body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return body

// }
