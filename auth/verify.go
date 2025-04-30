// verify authパッケージ
package auth

import (
	//"log/slog"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// ---- Global Variable

// ---- Package Global Variable

// ---- public function ----
func ParseJWTFromHeader(r *http.Request) (*Claims, error) {

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("no Authorization header")
	}

	// Bearer トークン形式
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("invalid Authorization header format")
	}

	tokenStr := parts[1]

	// 検証
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 型アサーション
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}