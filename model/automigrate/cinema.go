package automigrate

import "gorm.io/gorm"

// Cinema 影厅
type Cinema struct {
	gorm.Model
	Avatar string `gorm:"type:varchar(1000);not null"` // 电影院照片
	Name   string `gorm:"type;varchar(100)"`           // 电影院名字
	Rows   int    `gorm:"type:int;not null"`           // 座位行数
	cols   int    `gorm:"type:int;not null"`           // 座位列数
}
