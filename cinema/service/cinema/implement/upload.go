package implement

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"cinema/errs"
	"cinema/logs"
	"cinema/service/cinema"
	"cinema/service/cinema/input"
	"cinema/service/cinema/output"
)

func (impl *implementation) Upload(ctx context.Context, client cinema.HttpClienter, body *bytes.Buffer, writer *multipart.Writer, in *input.CinemaInput) (out *output.Cinema, err error) {
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
		bodyString := strings.Trim(string(bodyBytes), "\"")
		filters := []string{
			fmt.Sprintf("_id:eq:%v", in.ID),
		}

		update := bson.M{"image": bodyString, "updatedAt": time.Now()}

		err = impl.cinemaRepo.Update(ctx, filters, update)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
	}

	return out, nil
}
