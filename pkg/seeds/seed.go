package seeds

import (
	"gorm.io/gorm"
)

type Seed struct {
	Name string
	Run  func(orm *gorm.DB) error
}
