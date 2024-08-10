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

func InsertTireImages(ctx context.Context, images []string, id uint32) error {
	s := "["
	for _, v := range images {
		s += "'" + v + "',"
	}
	new := s[:len(s)-1] + "]"
	qry := fmt.Sprintf(query.InsertTireImage, new, id)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func GetTire(ctx context.Context, inspection_id uint32) (tire model.Tire, err error) {
	qr := fmt.Sprintf(query.GetTire, inspection_id)
	err = database.Conn.QueryRow(ctx, qr).ScanStruct(&tire)
	return
}

func GetTireImage(ctx context.Context, id uint32) ([]string, error) {
	qry := fmt.Sprintf(query.GetTireImage, id)
	var images []string
	rows, err := database.Conn.Query(ctx, qry)
	if err != nil {
		return images, err
	}
	for rows.Next() {
		var col1 []string
		if err := rows.Scan(&col1); err != nil {
			return images, err
		}
		images = append(images, col1...)
	}
	return images, err
}
