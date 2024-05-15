package automigrate

import "gorm.io/gorm"

type SeatStatus string

const (
	SeatLockStatus SeatStatus = "broken"
	SeatNormal     SeatStatus = "normal"
)

// Seat 座位信息
type Seat struct {
	gorm.Model
	CinemaID uint       // 影厅ID
	Cinema   Cinema     `gorm:"foreignKey:CinemaID;references:ID"` // 影厅
	Row      int        `gorm:"type:int;not null"`                 // 座位行
	Col      int        `gorm:"type:int;not null"`                 // 座位列
	Status   SeatStatus `gorm:"type:varchar(15)"`                  // 座位状态
}
