package auth

import (
	"blog/config"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Password string `json:"password"`
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string, password string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := []byte(config.Cfg.SecretKey)
	return token.SignedString(jwtSecret)
}
