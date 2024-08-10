package services

import (
	"context"
	"fmt"

	"github.com/ScoobieNoobie/Caterpillar/internal/database"
	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/query"
)

func CreateExterior(ctx context.Context, exterior model.GetExterior) error {
	qry := fmt.Sprintf(query.InsertExterior, fmt.Sprint(exterior.DamageToExterior), exterior.InspectionID, fmt.Sprint(exterior.OilLeak), exterior.ExteriorSummary)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func InsertExteriorImage(ctx context.Context, img string, id uint32) error {
	qry := fmt.Sprintf(query.InsertExteriorImage, img, id)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func InsertExteriorImages(ctx context.Context, imgs []string, id uint32) error {
	s := "["
	for _, v := range imgs {
		s += "'" + v + "',"
	}
	new := s[:len(s)-1] + "]"
	qry := fmt.Sprintf(query.InsertExteriorImages, new, id)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func GetExterior(ctx context.Context, id uint32) (ext model.Exterior, err error) {
	qry := fmt.Sprintf(query.GetExterior, id)
	err = database.Conn.QueryRow(ctx, qry).ScanStruct(&ext)
	return
}
