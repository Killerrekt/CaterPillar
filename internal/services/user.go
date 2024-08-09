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

func GetUser(ctx context.Context, employee_id string) (model.User, error) {
	var user model.User
	qry := fmt.Sprintf(query.GetUser, employee_id)
	err := database.Conn.QueryRow(ctx, qry).ScanStruct(&user)
	return user, err
}
