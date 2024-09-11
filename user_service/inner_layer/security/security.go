package security

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"errors"

	domainErrors "shop/user_service/inner_layer/domain/errors"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

const (
	Refresh = "refresh"
	Access  = "access"
)

type TokenClaims struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

// AppToken is a struct that contains the JWT token
type AppToken struct {
	Token          string    `json:"token"`
	TokenType      string    `json:"type"`
	ExpirationTime time.Time `json:"expitationTime"`
}

// TokenTypeKeyName is a map that contains the key name of the JWT in config.json
var TokenTypeKeyName = map[string]string{
	Access:  "Secure.JWTAccessSecure",
	Refresh: "Secure.JWTRefreshSecure",
}

// Structure likes in config.json
type SecureConfig struct {
	Secure struct {
		JWTAcessSecure     string `json:"JWTAcessSecure"`
		JWTRefreshSecure   string `json:"JWTRefreshSecure"`
		JWTAcessTimeMinute string `json:"JWTAcessTimeMinute"`
		JWTRefreshTimeHour string `json:"JWTRefreshTimeHour"`
	} `json:"Secure"`
}

func GenerateJWTToken(userID int, tokenType string) (appToken *AppToken, err error) {
	file, errRead := os.ReadFile("config.json")
	if err := errRead; err != nil {
		_ = fmt.Errorf("fatal error in config file: %s", err.Error())
	}

	var config SecureConfig

	err = json.Unmarshal(file, &config)
	if err != nil {
		return
	}

	JWTSecureKey := config.Secure.JWTAcessSecure
	JWTExpTime := config.Secure.JWTAcessTimeMinute

	if tokenType == Refresh {
		JWTSecureKey = config.Secure.JWTRefreshSecure
		JWTExpTime = config.Secure.JWTRefreshTimeHour
	}

	tokenTime, err := strconv.ParseInt(JWTExpTime, 10, 64)
	if err != nil {
		return
	}

	tokenTimeUnix := time.Duration(tokenTime)
	switch tokenType {
	case Access:
		tokenTimeUnix *= time.Minute
	case Refresh:
		tokenTimeUnix *= time.Hour
	default:
		err = errors.New("invalid token type")
	}

	if err != nil {
		return
	}

	tokenExpirationTime := time.Now().Add(tokenTimeUnix)

	claims := &TokenClaims{
		ID:   userID,
		Type: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(tokenExpirationTime),
		},
	}
	tokenWithNewClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := tokenWithNewClaims.SignedString([]byte(JWTSecureKey))
	if err != nil {
		return
	}

	appToken = &AppToken{
		Token:          tokenStr,
		TokenType:      tokenType,
		ExpirationTime: tokenExpirationTime,
	}

	return
}

func VerifyTokenAndGetClaims(token, tokenType string) (claims jwt.MapClaims, err error) {
	file, err := os.ReadFile("config.json")
	if err != nil {
		_ = fmt.Errorf("fatal error in config file: %s", err.Error())
		return
	}

	var config SecureConfig
	err = json.Unmarshal(file, &config)
	if err != nil {
		return
	}

	JWTRefreshSecureKey := config.Secure.JWTAcessSecure
	if tokenType != Access {
		JWTRefreshSecureKey = config.Secure.JWTRefreshSecure
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			errorString := fmt.Sprintf("wrong signing method %v", t.Header["alg"])
			return nil, &domainErrors.AppError{
				Err:  errors.New(errorString),
				Type: domainErrors.NotAuthenticated,
			}
		}

		return []byte(JWTRefreshSecureKey), nil
	})

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		if claims["type"] != tokenType {
			return nil, &domainErrors.AppError{
				Err:  errors.New("invalid token type"),
				Type: domainErrors.NotAuthenticated,
			}
		}

		var expTime = claims["exp"].(float64)
		if time.Now().Unix() > int64(expTime) {
			return nil, &domainErrors.AppError{
				Err:  errors.New("token expired"),
				Type: domainErrors.NotAuthenticated,
			}
		}

		return claims, nil
	}

	return
}

func IsFineCheckPasswordAndHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}

func GeneratePasswordHash(password string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	return
}
