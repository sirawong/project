package utils

import (
	qrcode "github.com/skip2/go-qrcode"
	"github.com/vincent-petithory/dataurl"
)

//go:generate mockery --name=QRCode
type QRCode interface {
	NewQRcode(text string) (string, error)
}

func NewQRcode(text string) (string, error) {
	qr, err := qrcode.Encode(text, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	url := dataurl.EncodeBytes(qr)
	return url, nil
}
