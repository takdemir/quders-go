package middleware

import (
	"database/pkg/controllers"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JWTWithConfig(SigningKey []byte, h *controllers.Handler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("x-api-token")
			if strings.TrimSpace(auth) == "" {
				return c.JSON(http.StatusForbidden, errors.New("invalid x-api-token header"))
			}
			token, _ := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return SigningKey, nil
			})
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && strings.TrimSpace(claims["email"].(string)) != "" {
				user, err := h.UserRepository.FindUserByEmail(claims["email"].(string))
				if err != nil {
					return c.JSON(http.StatusForbidden, errors.New("you are not authorized"))
				}
				c.Set("user", user)
				return next(c)
			}
			return c.JSON(http.StatusForbidden, errors.New("you are not authorized"))
		}
	}
}
