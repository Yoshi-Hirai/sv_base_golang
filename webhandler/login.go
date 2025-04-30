// login webhandlerパッケージ
package webhandler // パッケージ名はディレクトリ名と同じにする

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"sv_base/auth"
)

// ---- struct

// LoginRequestストラクト
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// ---- Global Variable

// ---- Package Global Variable

// ---- public function ----
func HandlerLogin(w http.ResponseWriter, r *http.Request) {

	slog.Info("Request", "form", r.Form, "path", r.URL.Path, "scheme", r.URL.Scheme, "url_long", r.Form["url_long"])
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// ここでパスワード認証（簡易的にスキップ）
	if req.Username != "user" || req.Password != "pass" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(req.Username)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}