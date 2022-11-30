package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

func ExtractWithClaims(tokenString string, signingKey []byte) (jwt.MapClaims, error) {
	var (
		jwtClaims jwt.MapClaims
	)
	if tokenString == "" {
		jwtClaims["role"] = "unauthorized"
		return jwtClaims, nil
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims, func(t *jwt.Token) (interface{}, error) {
		//checks weather parser method is same as we use
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwt.MapClaims)
	if err := claims.Valid(); err != nil && !ok {
		return nil, err
	}

	return *claims, nil
}
