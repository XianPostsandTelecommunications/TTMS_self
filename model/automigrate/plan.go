package automigrate

import (
	"gorm.io/gorm"
	"time"
)

// Plan 电影放映计划
type Plan struct {
	gorm.Model
	MovieID  uint      // 存储电影ID
	Movie    Movie     `gorm:"foreignKey:MovieID;references:ID"` // 关联的电影结构体
	CinemaID uint      // 影厅ID
	Cinema   Cinema    `gorm:"foreignKey:CinemaID;references:ID"` // 影厅结构体
	StartAT  time.Time // 电影放映开始时间
	EndAT    time.Time // 电影放映结束时间
	Price    float64   //电影票价
}
