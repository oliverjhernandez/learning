package api

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

// func JWTAuthentication(userStore db.UserStore) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		fmt.Println("-- JWT auth")
//
// 		token := c.Get("X-Api-Token")
// 		if len(token) == 0 {
// 			err := unauthorizedError(c)
// 			return err
// 		}
// 		claims, err := validateToken(token)
//		if err != nil {
// 			err := unauthorizedError(c)
// 			return err
// 		}
//
// 		// Checks token expiration
// 		fmt.Println(claims)
// 		expiresFloat := claims["expires"].(float64)
// 		expires := int64(expiresFloat)
// 		if time.Now().Unix() > expires {
// 			err := unauthorizedError(c)
// 			return err
// 		}
//
// 		userID := claims["id"].(float64)
//		user, err := userStore.GetUserByID(c.Context(), nil, int(userID))
//		if err != nil {
// 			err := unauthorizedError(c)
// 			return err
//
// 		}
// 		c.Context().SetUserValue("user", user)
//
// 		return c.Next()
// 	}
// }

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
