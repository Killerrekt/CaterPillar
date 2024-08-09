package database

import (
	"crypto/tls"
	"database/sql"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ScoobieNoobie/Caterpillar/config"
)

var Conn *sql.DB

func ConnectToDB() (err error) {
	Conn = clickhouse.OpenDB(&clickhouse.Options{
		Addr:     []string{config.ConfigData.ClickhouseURI},
		Protocol: clickhouse.Native,
		TLS:      &tls.Config{},
		Auth: clickhouse.Auth{
			Username: "default",
			Password: config.ConfigData.ClickhousePassword,
		},
	})
	return Conn.Ping()
}
