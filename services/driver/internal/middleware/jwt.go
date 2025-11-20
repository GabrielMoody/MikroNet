package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func GetJWTPayload(tokenString, secretKey string) (jwt.MapClaims, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

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

func ValidateDriverRole(c *fiber.Ctx) error {
	token := c.Get("Authorization")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  "Error",
			"message": "Unauthorized",
		})
	}

	claims, err := GetJWTPayload(token[7:], os.Getenv("JWT_SECRET"))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Error",
			"message": err.Error(),
		})
	}

	if claims["role"] != "driver" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":  "Error",
			"message": "Forbidden access",
		})
	}

	return c.Next()
}
