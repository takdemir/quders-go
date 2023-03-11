package middleware

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"notificaiton/pkg/controllers"
	"notificaiton/pkg/utils/httputils"
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
			//fmt.Println(ok, token.Valid, claims["email"].(string))
			if ok && strings.TrimSpace(claims["email"].(string)) != "" {

				requestBody := []byte(`{"email": "` + claims["email"].(string) + `"}`)

				url := fmt.Sprintf("%s/api/v1/user", h.DatabaseServiceHost)

				user, err := httputils.HttpPost(c, url, requestBody, "")

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
