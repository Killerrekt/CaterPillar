package services

import (
	"context"

	"github.com/ScoobieNoobie/Caterpillar/internal/database"
	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/query"
)

func InsertUser(ctx context.Context, name string, password string, employee_id uint32) error {
	err := database.Conn.Exec(ctx, query.InsertUser, name, password, employee_id)
	return err
}

func GetUser(ctx context.Context, employee_id uint32) (model.GetUser, error) {
	var user model.GetUser
	err := database.Conn.Select(ctx, &user, query.GetUser, employee_id)
	return user, err
}
