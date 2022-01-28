package utils

import (
	"errors"
	"fmt"
	"time"
	"user/config"
	"user/logs"

	jwt "github.com/dgrijalva/jwt-go"
)

//go:generate mockery --name=JWTService
type JWTService interface {
	GenerateToken(id string) (token *string, err error)
	ValidateToken(encodedToken string) (id string, err error)
}

type jwtServices struct {
	secretKey string
	issure    string
}

func NewJWT(appConfig *config.Config) JWTService {
	return &jwtServices{
		secretKey: appConfig.JWTSecret,
		issure:    "Thjk",
	}
}

func (service *jwtServices) GenerateToken(id string) (*string, error) {

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = id
	atClaims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return &t, nil
}

func (service *jwtServices) ValidateToken(encodedToken string) (id string, err error) {
	claimsToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})
	if err != nil {
		logs.Error(err)
		return "", err
	}

	claims, ok := claimsToken.Claims.(jwt.MapClaims)
	if ok && claimsToken.Valid {
		id, ok = claims["id"].(string)
		if !ok {
			return "", errors.New("cannot get id")
		}
	} else {
		logs.Error(err)
		return "", err
	}

	return id, nil
}
