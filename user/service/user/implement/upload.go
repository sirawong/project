package implement

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"user/errs"
	"user/logs"
	user "user/service/user"
	"user/service/user/input"
	"user/service/user/output"
)

func (impl *implementation) Upload(ctx context.Context, client user.HttpClienter, body *bytes.Buffer, writer *multipart.Writer, in *input.UserInput) (out *output.User, err error) {
	url := impl.config.PhotoUrl
	request, _ := http.NewRequest(http.MethodPost, url, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(request)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		filters := []string{
			fmt.Sprintf("_id:eq:%v", in.ID),
		}

		update := bson.M{"imageurl": bodyString, "updatedAt": time.Now()}

		err = impl.repo.Update(ctx, filters, update)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
	}

	return out, nil
}
