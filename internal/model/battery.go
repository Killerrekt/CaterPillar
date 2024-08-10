package model

type Battery struct {
	InspectionID           uint32   `json:"inspection_id" validate:"required"`
	BatteryMake            string   `json:"battery_make" validate:"required"`
	BatteryReplacementDate string   `json:"battery_replacement_date" validate:"required"`
	BatteryVoltage         uint32   `json:"battery_voltage" validate:"required"`
	BatteryWaterlevel      string   `json:"battery_water_level" validate:"required"`
	ConditionOfBattery     bool     `json:"condition_of_battery" validate:"required"`
	RustInBattery          bool     `json:"rust_in_battery" validate:"required"`
	BatterySummary         string   `json:"battery_in_summary" validate:"required"`
	BatteryImages          []string `json:"battery_images" validate:"required"`
}

type InsertBattery struct {
	InspectionID           uint32 `json:"inspection_id" validate:"required"`
	BatteryMake            string `json:"battery_make" validate:"required"`
	BatteryReplacementDate string `json:"battery_replacement_date" validate:"required"`
	BatteryVoltage         uint32 `json:"battery_voltage" validate:"required"`
	BatteryWaterlevel      string `json:"battery_water_level" validate:"required"`
	ConditionOfBattery     bool   `json:"condition_of_battery" validate:"required"`
	RustInBattery          bool   `json:"rust_in_battery" validate:"required"`
	BatterySummary         string `json:"battery_in_summary" validate:"required"`
}
