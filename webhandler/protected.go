// login webhandlerパッケージ
package webhandler // パッケージ名はディレクトリ名と同じにする

import (
	"encoding/json"
	//"log/slog"
	"net/http"

	"sv_base/auth"
)

// ---- Global Variable

// ---- Package Global Variable

// ---- public function ----
func HandlerProtected(w http.ResponseWriter, r *http.Request) {

	claims, err := auth.ParseJWTFromHeader(r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// トークンが有効なら、保護されたリソースを返す
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":  "Welcome!",
		"username": claims.Username,
	})
}