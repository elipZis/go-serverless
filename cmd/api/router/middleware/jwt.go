package middleware

import (
	"github.com/elipzis/go-serverless/api/handler"
	"github.com/elipzis/go-serverless/util"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected creates a Jwt Signed handler to secure routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(util.GetEnvOrDefault("JWT_SECRET", "verysecretsecret")),
		ErrorHandler: jwtError,
	})
}

// jwtError returns an error response in cases of bad/unauthorized requests
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return handler.SendError(c, fiber.StatusBadRequest, err)
	}
	return handler.SendErrorFromString(c, fiber.StatusUnauthorized, "Invalid or expired JWT")
}
