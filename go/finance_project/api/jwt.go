package api

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"finance/db"

	"github.com/golang-jwt/jwt/v5"

	"github.com/gofiber/fiber/v2"
)

func JWTAuthentication(userStore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fmt.Println("-- JWT auth")

		token := c.Get("X-Api-Token")
		if len(token) == 0 {
			return ErrUnauthorized()
		}
		claims, err := validateToken(token)
		if err != nil {
			return err
		}

		// Check token expiration
		expiresFloat := claims["expires"].(float64)
		expires := int64(expiresFloat)
		if time.Now().Unix() > expires {
			return NewError(http.StatusUnauthorized, "token expired")
		}

		userID := claims["id"].(float64)
		user, err := userStore.GetUserByID(c.Context(), nil, int(userID))
		if err != nil {
			return ErrUnauthorized()
		}
		c.Context().SetUserValue("user", user)

		return c.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Invalid signing method", token.Header["alg"])
			return nil, ErrUnauthorized()
		}

		secret := os.Getenv("JWT_SECRET")

		return []byte(secret), nil
	})
	if err != nil {
		fmt.Println("failed to parse jwt token", err)
		return nil, ErrUnauthorized()
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return nil, ErrUnauthorized()
	}

	return token.Claims.(jwt.MapClaims), nil
}
