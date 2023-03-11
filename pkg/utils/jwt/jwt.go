package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var JWTSecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateJWT(email string) string {
	err := godotenv.Load()
	if err != nil {
		return ""
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecretKey)
	return t
}
