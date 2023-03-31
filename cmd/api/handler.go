package main

import (
	"encoding/json"
)

func ResponseHandler(resp interface{}) ([]byte, error) {
	responseBody, err := json.Marshal(resp)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}
