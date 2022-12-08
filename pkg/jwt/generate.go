package jwt

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/AbduRohman-developer/cv-backend/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Tokens struct {
	Access           string
	Refresh          string
	AccessExpireTime int64
}

func New(credentials map[string]interface{}) (*Tokens, error) {
	//New access token with the given credentials
	accessToken, expire, err := newAccessToken(credentials)
	if err != nil {
		return nil, err
	}

	//New refresh token
	refreshToken, err := newRefreshToken()
	if err != nil {
		return nil, err
	}

	return &Tokens{
		Access:           accessToken,
		Refresh:          refreshToken,
		AccessExpireTime: expire,
	}, nil
}

func newAccessToken(credentials map[string]interface{}) (string, int64, error) {
	var (
		claims jwt.MapClaims
		cfg    = config.Get()
		expire int64
	)

	claims["iss"] = credentials["iss"]
	claims["role"] = credentials["role"]
	if cfg.Environment == "develop" {
		expire = time.Now().Add(time.Hour * 24 * 30).Unix()
	} else {
		expire = time.Now().Add(time.Hour * 24 * time.Duration(cfg.AccessKeyExpireDays)).Unix()
	}
	claims["exp"] = expire
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := t.SignedString(cfg.AccessSigningKey)
	if err != nil {
		return "", 0, err
	}

	return token, expire, nil
}

func newRefreshToken() (string, error) {
	var (
		hashSHA256 = sha256.New()
		cfg        = config.Get()
	)

	refresh := cfg.RefreshSigningKey + time.Now().String()

	if _, err := hashSHA256.Write([]byte(refresh)); err != nil {
		return "", err
	}

	token := hex.EncodeToString(hashSHA256.Sum(nil)) + fmt.Sprintf(".%d", time.Now().Add(time.Hour*24*time.Duration(cfg.RefreshKeyExpireDays)).Unix())

	return token, nil
}
