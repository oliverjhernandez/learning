package api

import (
	"fmt"
	"os"
	"time"

	"casita/internal/db"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("-- JWT auth")

		token := c.Get("X-Api-Token")
		if len(token) == 0 {
			unauthorizedError(c)
		}
		claims, err := validateToken(token)
		if err != nil {
			unauthorizedError(c)
		}

		// Check token expiration
		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)
		if time.Now().Unix() > expires {
			unauthorizedError(c)
		}

		userID := claims["id"].(float64)
		user, err := userStore.GetUserByID(c.Context(), nil, int(userID))
		if err != nil {
			unauthorizedError(c)
		}
		c.Context().SetUserValue("user", user)

		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	var err error
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, err
		}

		secret := os.Getenv("JWT_SECRET")

		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
