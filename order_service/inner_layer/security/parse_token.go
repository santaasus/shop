package security

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	root "shop"

	domainErrors "github.com/santaasus/errors-handler"
)

// Structure likes in config.json
type secureConfig struct {
	Secure struct {
		JWTAcessSecure string `json:"JWTAcessSecure"`
	} `json:"Secure"`
}

func ParseToken(token string) error {
	data, err := root.FileByName("config.json")
	if err != nil {
		err := domainErrors.ThrowAppErrorWith(domainErrors.InternalServerError)

		return err
	}

	var config secureConfig

	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	signature := []byte(config.Secure.JWTAcessSecure)

	if token == "" {
		err := &domainErrors.AppError{
			Err:  errors.New("token not provided"),
			Type: domainErrors.NotAuthenticated,
		}

		return err
	}

	var claims jwt.MapClaims
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return signature, nil
	})

	if err != nil {
		err := &domainErrors.AppError{
			Err:  errors.New("invalid token"),
			Type: domainErrors.ValidationError,
		}

		return err
	}

	return nil
}
