// pkg_webhandler WebHandlerパッケージ
package pkg_webhandler // パッケージ名はディレクトリ名と同じにする

import (
	"log/slog"
	"net/http"
)

// ---- Package Global Variable

// ---- public function ----
// Webhandler_init (public)Webhandlerシステムの初期化関数。handlerを登録し指定されたポートでサーバーを立ち上げる
func Webhandler_init(port string) error {
	slog.Info("WebHandler Open", "port", port)
	webhandlerErr := http.ListenAndServe(":"+port, nil)
	if webhandlerErr != nil {
		return webhandlerErr
	}

	return nil
}
