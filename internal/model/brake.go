package model

type GetBrake struct {
	InspectionID        uint32 `json:"inspection_id"`
	BrakeFluidLevel     string `json:"brake_fluid_level"`
	BrakeConditionFront string `json:"brake_condition_front"`
	BrakeConditionRear  string `json:"brake_condition_back"`
	EmergencyBrake      string `json:"emergency_break"`
	BrakeSummary        string `json:"brake_summary"`
}

type Brake struct {
	InspectionID        uint32   `json:"inspection_id"`
	BrakeFluidLevel     string   `json:"brake_fluid_level"`
	BrakeConditionFront string   `json:"brake_condition_front"`
	BrakeConditionRear  string   `json:"brake_condition_back"`
	EmergencyBrake      string   `json:"emergency_break"`
	BrakeSummary        string   `json:"brake_summary"`
	Images              []string `json:"images"`
}
