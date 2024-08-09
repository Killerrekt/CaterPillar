package model

type Tire struct {
	InspectionID           uint32   `json:"inspection_id" validate:"required"`
	TirePressureLeftFront  uint32   `json:"tire_pressure_left_front" validate:"required"`
	TirePressureRightFront uint32   `json:"tire_pressure_right_front" validate:"required"`
	TirePressureLeftRear   uint32   `json:"tire_pressure_left_rear" validate:"required"`
	TirePressureRightRear  uint32   `json:"tire_pressure_right_rear" validate:"required"`
	LeftFrontCondition     string   `json:"left_front_condition" validate:"required"`
	LeftRearCondition      string   `json:"left_rear_condition" validate:"required"`
	RightFrontCondition    string   `json:"right_front_condition" validate:"required"`
	RightRearCondition     string   `json:"right_rear_condition" validate:"required"`
	TireSummary            string   `json:"tire_summary" validate:"required"`
	Images                 []string `json:"images" validate:"required"`
}

type InsertTire struct {
	InspectionID           uint32 `json:"inspection_id" validate:"required"`
	TirePressureLeftFront  uint32 `json:"tire_pressure_left_front" validate:"required"`
	TirePressureRightFront uint32 `json:"tire_pressure_right_front" validate:"required"`
	TirePressureLeftRear   uint32 `json:"tire_pressure_left_rear" validate:"required"`
	TirePressureRightRear  uint32 `json:"tire_pressure_right_rear" validate:"required"`
	LeftFrontCondition     string `json:"left_front_condition" validate:"required"`
	LeftRearCondition      string `json:"left_rear_condition" validate:"required"`
	RightFrontCondition    string `json:"right_front_condition" validate:"required"`
	RightRearCondition     string `json:"right_rear_condition" validate:"required"`
	TireSummary            string `json:"tire_summary" validate:"required"`
}
