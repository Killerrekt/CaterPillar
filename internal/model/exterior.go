package model

type GetExterior struct {
	DamageToExterior bool   `json:"damage_to_exterior" validate:"required"`
	InspectionID     uint32 `json:"inspection_id" validate:"required"`
	OilLeak          bool   `json:"oil_leak" validate:"required"`
	ExteriorSummary  string `json:"external_summary" validate:"required"`
}

type Exterior struct {
	DamageToExterior bool     `json:"damage_to_exterior"`
	InspectionID     uint32   `json:"inspection_id"`
	OilLeak          bool     `json:"oil_leak"`
	ExteriorSummary  string   `json:"external_summary"`
	ExteriorImage    string   `json:"exterior_image"`
	Images           []string `json:"images"`
}
