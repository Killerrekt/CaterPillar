package middleware

import (
	"fmt"

	"github.com/ScoobieNoobie/Caterpillar/config"
	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func AuthenticateAndAuthorizeUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessSecret := config.ConfigData.AccessTokenSecret

		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Authorization header missing"})
		}

		accessToken := authHeader[len("Bearer "):]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(accessSecret), nil
		})

		fmt.Println(token)

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(model.Response{Message: "Not a valid token", Status: false})
		}

		employee_id, ok := claims["EmployeeID"].(string)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(model.Response{Message: "Failed to parse the claims", Status: false})
		}

		c.Locals("EmployeeID", employee_id)

		return c.Next()
	}
}
