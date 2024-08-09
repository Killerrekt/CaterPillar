package model

type User struct {
	Name       string `validate:"required" json:"name"`
	Password   string `validate:"required" json:"password"`
	EmployeeID string `validate:"required" json:"employee_id"`
}

type GetUser struct {
	Name       string `json:"name"`
	EmployeeID string `json:"employee_id"`
}
