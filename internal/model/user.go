package model

type User struct {
	Name       string `validate:"required" bson:"Name" json:"name"`
	Password   string `validate:"required" bson:"Password" json:"password"`
	EmployeeID string `validate:"required" bson:"EmployeeID" json:"employee_id"`
}
