package model

type GetEngine struct {
	InspectionID        uint32 `json:"inspection_id" validate:"required"`
	DamageEngine        bool   `json:"damage" validate:"required"`
	EngineOilCondition  string `json:"engine_oil_condition" validate:"required"`
	EngineOilColor      string `json:"engine_oil_color" validate:"required"`
	BrakeFluidCondition string `json:"brake_fluid_condition" validate:"required"`
	BrakeFluidColor     string `json:"brake_fluid_color" validate:"required"`
	OilLeak             bool   `json:"oil_leak" validate:"required"`
	Summary             string `json:"summary" validate:"required"`
}

type Engine struct {
	InspectionID        uint32   `json:"inspection_id" validate:"required"`
	DamageEngine        bool     `json:"damage" validate:"required"`
	EngineOilCondition  string   `json:"engine_oil_condition" validate:"required"`
	EngineOilColor      string   `json:"engine_oil_color" validate:"required"`
	BrakeFluidCondition string   `json:"brake_fluid_condition" validate:"required"`
	BrakeFluidColor     string   `json:"brake_fluid_color" validate:"required"`
	OilLeak             bool     `json:"oil_leak" validate:"required"`
	Summary             string   `json:"summary" validate:"required"`
	DmgEngineImg        string   `json:"damage_image"`
	Images              []string `json:"images"`
}
