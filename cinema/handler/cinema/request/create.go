package request

import (
	"bytes"
	"cinema/service/cinema/input"
	"encoding/json"
	"net/http"
)

func MakeCreateReq(input *input.CinemaInput) (req *http.Request, err error) {

	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("POST", "/cinemas", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	return req, nil
}

func MakeCreateReqInvalidJSON() (req *http.Request, err error) {

	jsonBytes := []byte("{{{}}}")

	req, err = http.NewRequest("POST", "/cinemas", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
