package jwt

import (
	"fmt"
	"github.com/AbduRohman-developer/cv-backend/config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type TokenMetaData struct {
	UserID     string
	Expires    int64
	Additional map[string]interface{}
}

func ExtractMetaData(c *gin.Context) (*TokenMetaData, error) {
	token, err := getToken(c)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		id, ok := claims["iss"].(string)
		if ok {
			if _, err := uuid.Parse(id); err != nil {
				return nil, err
			}
		}

		expire, ok := claims["exp"].(int64)
		if !ok {
			return nil, fmt.Errorf("cannot get expire date, %v", claims["exp"])
		}

		return &TokenMetaData{
			UserID:     id,
			Expires:    expire,
			Additional: claims,
		}, nil
	}

	return nil, fmt.Errorf("cannot parse claims, %v", claims)
}

func getToken(c *gin.Context) (*jwt.Token, error) {
	var (
		cfg = config.Get()
	)

	tokenString := c.GetHeader("Authorization")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.AccessSigningKey), nil
	})
	if err != nil {
		return nil, err
	}

	if err := verify(token); err != nil {
		return nil, err
	}

	return token, nil
}

func verify(token *jwt.Token) error {
	return token.Claims.Valid()
}
