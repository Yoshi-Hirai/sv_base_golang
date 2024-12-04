package main

import (
	"log/slog"

	"sv_base/config"
	"sv_base/log"
	"sv_base/test"
	"sv_base/webhandler"
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
	test.TestJsonFileReadWrite()
	// csvファイル読み書きテスト
	test.TestCsvFileReadWrite()

	// Redisへの接続テスト
	test.TestRedisConnect()

	// PSQLへの接続テスト
	test.TestPostgresConnect()

	// 指定ポートでサーバーを立ち上げる
	errWeb := webhandler.WebHandlerInit("8080")
	if errWeb != nil {
		slog.Info("Webhandler Error", "err", errWeb)
		return
	}

}
