package services

import (
	"context"
	"fmt"

	"github.com/ScoobieNoobie/Caterpillar/internal/database"
	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/query"
)

func InsertTire(ctx context.Context, tire model.InsertTire) error {
	qry := fmt.Sprintf(query.InsertTire, tire.InspectionID,
		tire.TirePressureLeftFront,
		tire.TirePressureRightFront,
		tire.TirePressureLeftRear,
		tire.TirePressureRightRear,
		tire.LeftFrontCondition,
		tire.LeftRearCondition,
		tire.RightFrontCondition,
		tire.RightRearCondition,
		tire.TireSummary)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func InsertTireImages(ctx context.Context, images []string) error {
	qry := fmt.Sprintf(query.InsertTireImage, images)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func GetTire(ctx context.Context, inspection_id uint32) (tire model.Tire, err error) {
	qr := fmt.Sprintf(query.GetTire, inspection_id)
	err = database.Conn.QueryRow(ctx, qr).ScanStruct(&tire)
	return
}
