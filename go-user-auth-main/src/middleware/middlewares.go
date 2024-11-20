package middlewares

import (
	"go-user-auth/config"
	"go-user-auth/utils"

	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No token provided"})
		}
		claims, err := utils.VerifyJWTToken(authHeader, config.KeyId)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}
		c.Locals("ClientId", claims["ClientId"])
		return c.Next()
	}
}
