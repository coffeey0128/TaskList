package models

type Task struct {
	Id     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name;NOT NULL;type:varchar(255)" json:"name"`
	Status int    `gorm:"column:status;default:0" json:"status"` // 0: Incompleted, 1: Completed
}
