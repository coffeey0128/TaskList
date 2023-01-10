package models

type Task struct {
	Id     int64  `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	Status int64  `gorm:"column:status" json:"status"`
}
