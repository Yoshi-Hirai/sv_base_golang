// jwt authパッケージ
package auth

import (
	//"log/slog"
	"time"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// ---- struct

type Claims struct {
	Username string `json:"username"`	// カスタムClaim
	jwt.RegisteredClaims				// 標準Claim (exp, iat, etc.)
}

// ---- Global Variable

// ---- Package Global Variable
var jwtKey = []byte("your-secret-key") // [TODO] あとでconfigから読めるようにしてもOK

// ---- public function ----
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

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 1. AuthorizationヘッダからJWTを取り出す
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        // 2. トークンを "Bearer ..." 形式から抽出
        parts := strings.Split(tokenStr, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Invalid token format", http.StatusUnauthorized)
            return
        }
        tokenStr = parts[1]

        // 3. JWTを検証
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // 4. 検証OKなら次の処理へ
        next.ServeHTTP(w, r)
    })
}