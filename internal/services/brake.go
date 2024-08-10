package services

import (
	"context"
	"fmt"

	"github.com/ScoobieNoobie/Caterpillar/internal/database"
	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/query"
)

func CreateBrake(ctx context.Context, Brake model.GetBrake) error {
	qry := fmt.Sprintf(query.InsertBrake,
		Brake.InspectionID,
		Brake.BrakeFluidLevel,
		Brake.BrakeConditionFront,
		Brake.BrakeConditionRear,
		Brake.EmergencyBrake,
		Brake.BrakeSummary)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func InsertBrakeImages(ctx context.Context, imgs []string, id uint32) error {
	s := "["
	for _, v := range imgs {
		s += "'" + v + "',"
	}
	new := s[:len(s)-1] + "]"
	qry := fmt.Sprintf(query.InsertBrakeImages, new, id)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func GetBrake(ctx context.Context, id uint32) (ext model.Brake, err error) {
	qry := fmt.Sprintf(query.GetBrake, id)
	err = database.Conn.QueryRow(ctx, qry).ScanStruct(&ext)
	return
}
