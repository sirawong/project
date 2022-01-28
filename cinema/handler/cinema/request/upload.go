package request

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
)

func MakeUploadReq(id string, file *bytes.Buffer, writer *multipart.Writer) (req *http.Request, err error) {

	req, err = http.NewRequest("POST", fmt.Sprintf("/photo/%s", id), file)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}
