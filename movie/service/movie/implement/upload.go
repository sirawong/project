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

	"movie/errs"
	"movie/logs"
	movie "movie/service/movie"
	"movie/service/movie/input"
	"movie/service/movie/output"
)

func (impl *implementation) Upload(ctx context.Context, client movie.HttpClienter, body *bytes.Buffer, writer *multipart.Writer, in *input.MovieInput) (out *output.Movie, err error) {
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

		err = impl.repo.Update(ctx, filters, update)
		if err != nil {
			logs.Error(err)
			return nil, errs.NewUnexpectedError()
		}
	}

	return out, nil
}
