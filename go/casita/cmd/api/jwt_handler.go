package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"casita/db"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("-- JWT auth")

		token := c.Get("X-Api-Token")
		if len(token) == 0 {
			return NewError(http.StatusUnauthorized, UNAUTHORIZED)
		}
		claims, err := validateToken(token)
		if err != nil {
			return NewError(http.StatusUnauthorized, INVALID_TOKEN)
		}

		// Check token expiration
		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)
		if time.Now().Unix() > expires {
			return NewError(http.StatusUnauthorized, TOKEN_EXPIRED)
		}

		userID := claims["id"].(float64)
		user, err := userStore.GetUserByID(c.Context(), nil, int(userID))
		if err != nil {
			return NewError(http.StatusUnauthorized, UNAUTHORIZED)
		}
		c.Context().SetUserValue("user", user)

		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, NewError(http.StatusUnauthorized, UNAUTHORIZED)
		}

		secret := os.Getenv("JWT_SECRET")

		return []byte(secret), nil
	})
	if err != nil {
		return nil, NewError(http.StatusUnauthorized, INVALID_TOKEN)
	}

	if !token.Valid {
		return nil, NewError(http.StatusUnauthorized, TOKEN_EXPIRED)
	}

	return token.Claims.(jwt.MapClaims), nil
}
