// redis Reidsシステムパッケージ
package redis // パッケージ名はディレクトリ名と同じにする

import (
	"context"
	"log/slog"

	"github.com/redis/go-redis/v9"
)

// ---- Global Variable

// ---- Package Global Variable
var rdb *redis.Client

//---- public function ----

// Redis_init (public)Redisシステムの初期化関数。
func RedisInit() error {

	var ctx = context.Background()

	// [ToDo] 接続情報を外出しする
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 1000,
	})

	pong, errRedis := rdb.Ping(ctx).Result()
	if errRedis != nil {
		slog.Info("Reids Error", "err", errRedis, "pong", pong)
		return errRedis
	}
	slog.Info("Redis Open Success.")

	//rdb.Set(ctx, "mykey1", "hoge", 0)
	ret, err := rdb.Get(ctx, "test-key").Result()
	if err != nil {
		slog.Info("Error: ", "err", err)
		return err
	}
	slog.Info("Redis Get", "test-key", ret)

	return nil
}

//---- private function ----
