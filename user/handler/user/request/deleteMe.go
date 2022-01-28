package request

import (
	"net/http"
)

func MakeDeleteMeReq() (req *http.Request, err error) {

	req, err = http.NewRequest("DELETE", "/users/me", nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
