package models

 
type (
	BaseTable struct {
		Id int `gorm:"primarykey"`
		Name string `gorm:"not null"`
	}
)