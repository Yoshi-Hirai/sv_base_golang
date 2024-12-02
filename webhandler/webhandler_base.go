// webhandler webhandlerパッケージ
package webhandler // パッケージ名はディレクトリ名と同じにする

import (
	"log/slog"
	"net/http"
)

// ---- Package Global Variable

// ---- public function ----
// WebHandlerInit (public)Webhandlerシステムの初期化関数。handlerを登録し指定されたポートでサーバーを立ち上げる
func WebHandlerInit(port string) error {
	slog.Info("WebHandler Open", "port", port)
	webhandlerErr := http.ListenAndServe(":"+port, nil)
	if webhandlerErr != nil {
		return webhandlerErr
	}

	return nil
}
