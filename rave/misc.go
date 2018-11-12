// utils and helper functions

package rave

import (
	"net/http"
	"bytes"
	"strings"
	"fmt"
	"encoding/json"
	// "io/ioutil"
	"runtime"
	
	// "github.com/antonholmquist/jason"
)
// Converts map[string]interface{} to JSON
func mapToJSON(mapData map[string]interface{}) []byte {
	jsonBytes, err := json.Marshal(mapData)
	if err != nil {
		panic(err)
	}

	return jsonBytes
}

// Checks if all required parameters are present
func checkRequiredParameters(params map[string]interface{}, keys []string) error {
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
func makePostRequest(data map[string]interface{}, URL string) (map[string]interface{}) {
    
   
    postData, err := json.Marshal(data)
    if err != nil {
        fmt.Println(err)
    }
    
	resp, err := http.Post(URL, "application/json", bytes.NewBuffer(postData))
    
    var result map[string]interface{}
    json.NewDecoder(resp.Body).Decode(&result)

    return result

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