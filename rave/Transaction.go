package rave

import "errors"

// import (
// 	"time"
// )

type Transaction map[string]interface{}

func TransactionGetID(id string, secretKey string) (response Transaction, err error) {
	url := rave.GetEndpoint("transaction", "get.id")
	res, err := GetRequest(url, map[string]string{":id": id}, nil, rave.GetSecretKey())
	if err != nil {
		return nil, err
	}
	status, _ := res["status"].(string)
	if status != "success" {
		return res, errors.New("api call not successfull")
	}
	response, ok := res["data"].(map[string]interface{})
	if !ok {
		return res, errors.New("invalid data returned")
	}
	return response, nil
}

func TransactionGetMany(query map[string]string, secretKey string) (response []Transaction, data interface{}, total float64, currentPage float64, totalPages float64, err error) {
	url := rave.GetEndpoint("transaction", "get.many")
	res, err := GetRequest(url, nil, query, rave.GetSecretKey())
	if err != nil {
		return nil, nil, 0, 0, 0, err
	}
	status, _ := res["status"].(string)
	if status != "success" {
		return nil, res, 0, 0, 0, errors.New("api call not successfull")
	}
	meta, ok := res["meta"].(map[string]interface{})
	if !ok {
		return nil, res, 0, 0, 0, errors.New("invalid data returned for field 'meta'")
	}

	pageInfo, ok := meta["page_info"].(map[string]interface{})
	if !ok {
		return nil, res, 0, 0, 0, errors.New("invalid data returned for field 'pageInfo'")
	}

	total, _ = pageInfo["total"].(float64)
	currentPage, _ = pageInfo["current_page"].(float64)
	totalPages, _ = pageInfo["total_pages"].(float64)

	dat, ok := res["data"].([]map[string]interface{})
	if !ok {
		return nil, res, 0, 0, 0, errors.New("invalid data returned for field 'data'")
	}

	for _, v := range dat {
		response = append(response, v)
	}

	return
}
