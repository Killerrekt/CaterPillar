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
	return
}
