package services

import (
	"context"
	"fmt"

	"github.com/ScoobieNoobie/Caterpillar/internal/database"
	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/query"
)

func GetHeader(ctx context.Context, employee_id string) (header model.GetHeader, err error) {
	qry := fmt.Sprintf(query.GetHeader, employee_id)
	err = database.Conn.QueryRow(ctx, qry).ScanStruct(&header)
	return
}

func InsertHeader(ctx context.Context, header model.CreateHeader) (err error) {
	qry := fmt.Sprintf(query.InsertHeader,
		header.TruckSerialNumber,
		header.TruckModel,
		header.InspectionID,
		header.InspectorName,
		header.InspectorEmployeeID,
		header.Location,
		header.ClientID,
		header.ClientName,
		fmt.Sprint(header.Service))
	err = database.Conn.Exec(ctx, qry)
	return
}
