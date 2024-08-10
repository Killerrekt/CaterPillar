package model

import "time"

type Header struct {
	TruckSerialNumber   string    `json:"truck_serial_number"`
	TruckModel          string    `json:"truck_model"`
	InspectionID        uint32    `json:"inspection_id"`
	InspectorName       string    `json:"inspector_name"`
	InspectorEmployeeID string    `json:"inspector_employee_id"`
	DateAndTime         time.Time `json:"date_and_time"`
	Location            string    `json:"location"`
	ClientName          string    `json:"client_name"`
	ClientID            string    `json:"client_id"`
	Service             bool      `json:"sevice_status"`
}

type GetHeader struct {
	TruckSerialNumber   string    `json:"truck_serial_number"`
	TruckModel          string    `json:"truck_model"`
	InspectionID        uint32    `json:"inspection_id"`
	InspectorName       string    `json:"inspector_name"`
	InspectorEmployeeID string    `json:"inspector_employee_id"`
	DateAndTime         time.Time `json:"date_and_time"`
	Location            string    `json:"location"`
	ClientName          string    `json:"client_name"`
	ClientID            string    `json:"client_id"`
	Service             bool      `json:"sevice_status"`
	Image               string    `json:"string"`
}

type CreateHeader struct {
	TruckSerialNumber   string `json:"truck_serial_number" validate:"required"`
	TruckModel          string `json:"truck_model" validate:"required"`
	InspectionID        uint32 `json:"inspection_id" validate:"required"`
	InspectorName       string `json:"inspector_name" validate:"required"`
	InspectorEmployeeID string `json:"-"`
	Location            string `json:"location" validate:"required"`
	ClientName          string `json:"client_name" validate:"required"`
	ClientID            string `json:"client_id" validate:"required"`
	Service             bool   `json:"service_status"`
}
