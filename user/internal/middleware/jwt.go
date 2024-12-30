package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func GetJWTPayload(tokenString, secretKey string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	// Extract the claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid JWT claims")
	}
}