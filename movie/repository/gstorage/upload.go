package gstorage

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"time"
)

func (st *storageRepo) Upload(ctx context.Context, filename string, file multipart.File) (urlPath string, err error) {

	sw := st.st.Object(fmt.Sprintf("images/%v-%v", time.Now().Unix(), filename)).NewWriter(ctx)

	if _, err := io.Copy(sw, file); err != nil {
		return "", err
	}

	if err := sw.Close(); err != nil {
		return "", err
	}

	u, err := url.Parse("https://storage.googleapis.com/" + st.appConfig.BukgetName + "/" + sw.Attrs().Name)
	if err != nil {
		return "", err
	}

	return u.String(), nil
}
