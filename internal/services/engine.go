package services

import (
	"context"
	"fmt"

	"github.com/ScoobieNoobie/Caterpillar/internal/database"
	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/query"
)

func CreateEngine(ctx context.Context, Engine model.GetEngine) error {
	qry := fmt.Sprintf(query.InsertEngine,
		Engine.InspectionID,
		fmt.Sprint(Engine.DamageEngine),
		Engine.EngineOilCondition,
		Engine.EngineOilColor,
		Engine.BrakeFluidCondition,
		Engine.BrakeFluidColor,
		fmt.Sprint(Engine.OilLeak),
		Engine.Summary)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func InsertEngineImages(ctx context.Context, imgs []string, id uint32) error {
	s := "["
	for _, v := range imgs {
		s += "'" + v + "',"
	}
	new := s[:len(s)-1] + "]"
	qry := fmt.Sprintf(query.InsertEngineImages, new, id)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func InsertEngineImage(ctx context.Context, img string, id uint32) error {
	qry := fmt.Sprintf(query.InsertEngineImage, img, id)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func GetEngine(ctx context.Context, id uint32) (ext model.Engine, err error) {
	qry := fmt.Sprintf(query.GetEngine, id)
	err = database.Conn.QueryRow(ctx, qry).ScanStruct(&ext)
	return
}
