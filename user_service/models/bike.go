package models

// This is bike struct
type Bike struct {
	ID   string `json:"id" gorm:"column:id;primarykey"`
	Name string `json:"name" gorm:"column:name"`
}
