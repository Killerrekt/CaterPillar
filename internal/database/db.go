package database

import (
	"crypto/tls"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/ScoobieNoobie/Caterpillar/config"
)

var Conn driver.Conn

func ConnectToDB() (err error) {
	Conn, err = clickhouse.Open(&clickhouse.Options{
		Addr: []string{config.ConfigData.ClickhouseURI},
		Auth: clickhouse.Auth{
			Database: config.ConfigData.ClickhouseDatabase,
			Username: "default",
			Password: config.ConfigData.ClickhousePassword,
		},
		Debug: true,
		Debugf: func(format string, v ...any) {
			fmt.Printf(format+"\n", v...)
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		TLS: &tls.Config{
			InsecureSkipVerify: true,
		},
		DialTimeout:          time.Second * 30,
		MaxOpenConns:         5,
		MaxIdleConns:         5,
		ConnMaxLifetime:      time.Duration(10) * time.Minute,
		ConnOpenStrategy:     clickhouse.ConnOpenInOrder,
		BlockBufferSize:      10,
		MaxCompressionBuffer: 10240,
	})
	if err != nil {
		return
	}
	defer Conn.Close()
	return
}
