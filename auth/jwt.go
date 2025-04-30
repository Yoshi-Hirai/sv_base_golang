package auth

import (
	//"log/slog"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your-secret-key") // [TODO] あとでconfigから読めるようにしてもOK

type Claims struct {
	Username string `json:"username"`	// カスタムClaim
	jwt.RegisteredClaims				// 標準Claim (exp, iat, etc.)
}

func GenerateJWT(username string) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),	// 1時間有効
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}