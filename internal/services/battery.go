package services

import (
	"context"
	"fmt"

	"github.com/ScoobieNoobie/Caterpillar/internal/database"
	"github.com/ScoobieNoobie/Caterpillar/internal/model"
	"github.com/ScoobieNoobie/Caterpillar/internal/query"
)

func InsertBattery(ctx context.Context, battery model.InsertBattery) error {
	qry := fmt.Sprintf(query.InsertBattery, battery.InspectionID,
		battery.BatteryMake,
		battery.BatteryReplacementDate,
		battery.BatteryVoltage,
		battery.BatteryWaterlevel,
		fmt.Sprint(battery.ConditionOfBattery),
		fmt.Sprint(battery.RustInBattery),
		battery.BatterySummary)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func InsertBatteryImages(ctx context.Context, images []string) error {
	qry := fmt.Sprintf(query.InsertBatteryImages, images)
	err := database.Conn.Exec(ctx, qry)
	return err
}

func GetBattery(ctx context.Context, inspection_id uint32) (tire model.Battery, err error) {
	qr := fmt.Sprintf(query.GetBattery, inspection_id)
	err = database.Conn.QueryRow(ctx, qr).ScanStruct(&tire)
	return
}
