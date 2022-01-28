package request

import (
	"bytes"
	"encoding/json"
	"net/http"

	"user/entities"
)

func MakeAllReq(opt *entities.PageOption) (req *http.Request, err error) {

	jsonBytes, err := json.Marshal(opt)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("GET", "/users", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func MakeAllReqInvalidJSON() (req *http.Request, err error) {

	jsonBytes := []byte("{{{}}}")

	req, err = http.NewRequest("GET", "/users", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
