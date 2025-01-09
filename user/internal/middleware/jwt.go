package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func GetJWTPayload(tokenString, secretKey string) (jwt.MapClaims, error) {
	if tokenString == "" {
		return nil, fmt.Errorf("token is empty")
	}

	token, err := jwt.Parse(tokenString[7:], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid JWT claims")
	}
}

func ValidateUserRole(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	payload, err := GetJWTPayload(token, os.Getenv("JWT_SECRET"))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "error",
			"message": "Unauthorized",
		})
	}

	if payload["role"] != "user" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "error",
			"message": "forbidden access",
		})
	}

	return c.Next()
}
