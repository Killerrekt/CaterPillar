package database

import (
	"context"

	"github.com/ScoobieNoobie/Caterpillar/internal/query"
)

func RunMigration(ctx context.Context) (err error) {
	err = Conn.Exec(ctx, query.CreateUserTable)
	if err != nil {
		return
	}
	err = Conn.Exec(ctx, query.CreateHeaderTable)
	if err != nil {
		return
	}
	err = Conn.Exec(ctx, query.CreateTireTable)
	if err != nil {
		return
	}
	err = Conn.Exec(ctx, query.CreateBatteryTable)
	if err != nil {
		return
	}
	err = Conn.Exec(ctx, query.CreateExteriorTable)
	if err != nil {
		return
	}
	err = Conn.Exec(ctx, query.CreateBrakeTable)
	if err != nil {
		return
	}
	err = Conn.Exec(ctx, query.CreateEngineTable)
	if err != nil {
		return
	}
	return
}
