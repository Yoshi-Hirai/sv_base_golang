// db DBシステムパッケージ
package db // パッケージ名はディレクトリ名と同じにする

import (
	"database/sql"
	"log/slog"
	"sv_funapp/config"

	_ "github.com/go-sql-driver/mysql" // for mysql
	_ "github.com/lib/pq"              // for postgres
)

// ---- Global Variable

// ---- Package Global Variable
var db *sql.DB

//---- public function ----

// DbBaseInit (public)DataBaseシステムの初期化関数。
func DbBaseInit() error {

	// 接続情報をConfigから取得
	dbProperty := config.GetConfigInformation()
	var errOpen, errPing error
	driverName := dbProperty.DbDriver
	//dataSourceName := "host=127.0.0.1 port=5432 user=user dbname=mydatabase password=password sslmode=disable"
	dataSourceName := "host=" + dbProperty.DbHost + " port=" + dbProperty.DbPort + " user=" + dbProperty.DbUser + " dbname=" + dbProperty.DbName + " password=" + dbProperty.DbPasswd + " sslmode=" + dbProperty.DbSslMode
	/*
		c := mysql.Config{
			DBName:               dbProperty.DbName,
			User:                 dbProperty.DbUser,
			Passwd:               dbProperty.DbPasswd,
			Addr:                 dbProperty.DbHost,
			Net:                  dbProperty.DbNet,
			ParseTime:            dbProperty.DbParseTime,
			AllowNativePasswords: dbProperty.DbAllowNativePasswords,
			Collation:            dbProperty.DbCollation,
		}
		dataSourceName := c.FormatDSN()
	*/
	slog.Info("Database Open", "driver", driverName, "dsn", dataSourceName)

	db, errOpen = sql.Open(driverName, dataSourceName)
	if errOpen != nil {
		return errOpen
	}
	errPing = db.Ping()
	if errPing != nil {
		return errPing
	}

	slog.Info("DB Initialize.", "db", db)
	return nil
}

// 現在の日付と時刻を文字列で取得する
func GetNow() (string, error) {

	// 最終更新日
	sqlStr := "SELECT NOW()"
	slog.Info(sqlStr)
	rows, errQuery := db.Query(sqlStr)
	if errQuery != nil {
		return "", errQuery
	}

	var result string
	for rows.Next() {
		if errScan := rows.Scan(&result); errScan != nil {
			return "", errScan
		}
	}

	return result, nil
}

//---- private function ----
