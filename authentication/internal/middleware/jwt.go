package middleware

import (
	"github.com/GabrielMoody/MikroNet/authentication/internal/helper"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware() fiber.Handler {
	v := helper.LoadEnv()
	return jwtware.New(jwtware.Config{SigningKey: jwtware.SigningKey{Key: v.GetString("JWT_SECRET")}})
}
