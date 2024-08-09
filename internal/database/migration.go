package database

import (
	"context"
	"time"

	"github.com/ScoobieNoobie/Caterpillar/internal/query"
)

func RunMigration(ctx context.Context) (err error) {
	time.Sleep(2 * time.Second)
	err = Conn.Exec(ctx, query.CreateUserTable)
	if err != nil {
		return
	}
	return
}
