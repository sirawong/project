package request

import (
	"net/http"
)

func MakeGetMeReq() (req *http.Request, err error) {

	req, err = http.NewRequest("GET", "/users/me", nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
