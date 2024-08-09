package services

import (
	"context"
	"fmt"

	"github.com/ScoobieNoobie/Caterpillar/internal/database"
	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/query"
)

func InsertUser(ctx context.Context, user model.User) error {
	qry := fmt.Sprintf(query.InsertUser, user.Name, user.Password, user.EmployeeID)
	fmt.Println(qry)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func GetUser(ctx context.Context, employee_id string) (model.GetUser, error) {
	var user model.GetUser
	err := database.Conn.Select(ctx, &user, query.GetUser, employee_id)
	return user, err
}
