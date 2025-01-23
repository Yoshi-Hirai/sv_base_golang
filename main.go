package main

import (
	"log/slog"

	"sv_funapp_websocket/config"
	"sv_funapp_websocket/log"
	"sv_funapp_websocket/webhandler"
)

// ---- Global Variable
var IsDebug bool = true

// ---- Package Global Variable

// ---- public function ----

// ---- private function ----
// 初期化処理関数
func initialize() bool {

	// ロガー生成
	jsonLogger := log.GetInstance()
	jsonLogger.Info("main")

	// Config読み込み
	errConfig := config.ReadConfigInformation()
	if errConfig != nil {
		slog.Error("Read Config Failed")
		jsonLogger.Error("Read Config Failed", slog.String("error", errConfig.Error()))
		return false
	}
	return true
}

//

func main() {

	isSuccess := initialize()
	if isSuccess == false {
		slog.Error("Initialize Failed")
		return
	}

	// jsonファイル読み書きテスト
	//test.TestJsonFileReadWrite()
	// csvファイル読み書きテスト
	//test.TestCsvFileReadWrite()

	// Redisへの接続テスト
	//test.TestRedisConnect()
	// PSQLへの接続テスト
	//test.TestPostgresConnect()

	// convertパッケージに用意された関数を実行するサンプル処理
	//test.TestConvertPackage()

	webErr := webhandler.WebHandlerInit("8080")
	if webErr != nil {
		slog.Info("Webhandler Error", "err", webErr)
		return
	}
}
