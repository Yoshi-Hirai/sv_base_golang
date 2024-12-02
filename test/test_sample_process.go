// test テストパッケージ
package test // パッケージ名はディレクトリ名と同じにする

import (
	"log/slog"
	"sv_base/db"
	"sv_base/fileio"
	"sv_base/redis"
)

// jsonサンプルストラクト(TestJsonFileReadWriteで利用)
type MasterTokenUse struct {
	Header  Header    `json:"header"`
	Payload []Payload `json:"payload"`
}
type Header struct {
	ContentType string `json:"content_type"`
}
type Payload struct {
	TokenId        string `json:"token_id"`
	PlaceToConsume string `json:"place_to_consume"`
	Version        string `json:"version"`
}

// ---- Global Variable

// ---- Package Global Variable

//---- public function ----

// PostgresSQLに接続するサンプル処理
func TestPostgresConnect() {

	errDb := db.DbBaseInit()
	if errDb != nil {
		slog.Info("DB Error", "err", errDb)
		return
	}

	nowString, errQ := db.GetNow()
	if errQ != nil {
		slog.Info("DB Error", "err", errQ)
	}
	slog.Info("NOW", "time", nowString)

}

// Redisに接続するサンプル処理
func TestRedisConnect() {

	errRedis := redis.RedisInit()
	if errRedis != nil {
		slog.Info("REDIS Error", "err", errRedis)
		return
	}
}

// Jsonファイルを読み書きするサンプル処理
func TestJsonFileReadWrite() {

	// ファイルからJsonを読む
	var master_token_use_json MasterTokenUse
	err := fileio.FileIoJsonRead("data/master_token_use_wobom.json", &master_token_use_json)
	if err != nil {
		slog.Error("Read Json Failed", slog.String("error", err.Error()))
		return
	}
	slog.Info("Read Json", "data", slog.Any("data", master_token_use_json))

	// json配列に1要素追加しファイルに書く
	add_payload := Payload{"100", "10000", "1.09.8000"}
	master_token_use_json.Payload = append(master_token_use_json.Payload, add_payload)
	err = fileio.FileIoJsonWrite("data/master_token_use_rewrite.json", master_token_use_json, false)
	if err != nil {
		slog.Error("Write Json Failed", slog.String("error", err.Error()))
		return
	}

	slog.Info("TestJsonFileReadWrite", "Result", "Success")
}

// csvファイルを読み書きするサンプル処理
func TestCsvFileReadWrite() {

	cont, err := fileio.FileIoCsvRead("data/14KANAGA.csv")
	if err != nil {
		slog.Error("Read Csv Failed", slog.String("error", err.Error()))
		return
	}
	// [][]stringなのでループする
	/*
		for _, v := range cont {
			fmt.Println(v)
		}
	*/
	err = fileio.FileIoCsvWrite("data/14KANAGA_Rewrite.csv", cont, true)
	if err != nil {
		slog.Error("Write Csv Failed", slog.String("error", err.Error()))
		return
	}
}

//---- private function ----
