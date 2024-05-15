package automigrate

import "gorm.io/gorm"

type ManagerMovie struct {
	gorm.Model
	MovieID uint
	UserID  uint
	Content string `gorm:"type:varchar(1000)"`
}
