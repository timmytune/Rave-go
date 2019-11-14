package rave

// import (
// 	"strconv"
// )

type BVN struct {
	Rave
}

// type Bvn interface {
// 	Withdraw(data string) (error error, response map[string]interface{})
// }

type BvnData struct {
	BvnNumber   string `json:"bvn"`
	Seckey string `json:"seckey"`
}






func (b BVN) Bvn(data BvnData) (error error, response map[string]interface{}) {
	queryParam := map[string]string{
	
		"bvn":     data.BvnNumber,
		"seckey": b.GetSecretKey(),
		
	
	}
	url := b.GetBaseURL() + b.GetEndpoint("bvn", "bvnverification")
	
	err, response := MakeGetRequest(url, queryParam)
	if err != nil {
		return err, noresponse
	}
	return nil, response
}
