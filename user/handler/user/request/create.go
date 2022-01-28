package request

import (
	"bytes"
	"encoding/json"
	"net/http"
	"user/service/user/input"
)

func MakeCreateReq(input *input.UserInput) (req *http.Request, err error) {

	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func MakeCreateReqInvalidJSON() (req *http.Request, err error) {

	jsonBytes := []byte("{{{}}}")

	req, err = http.NewRequest("POST", "/users", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
