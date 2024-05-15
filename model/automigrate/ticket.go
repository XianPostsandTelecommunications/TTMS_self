package automigrate

import (
	"gorm.io/gorm"
	"time"
)

type TicketStatus string

const (
	LockStatus    TicketStatus = "lock"     // 锁定
	ForSaleStatus TicketStatus = "for_sale" // 待售
	SaledStatus   TicketStatus = "saled"    //已售
)

// Ticket 票务信息
type Ticket struct {
	gorm.Model
	PlanID       uint
	Plan         Plan `gorm:"foreignKey:PlanID;references:ID"` // 放映计划
	UserID       uint `gorm:"index"`
	SeatID       uint
	Seat         Seat         `gorm:"foreignKey:SeatID;references:ID"`
	Price        float64      `gorm:"type:float;not null"`
	TicketStatus TicketStatus `gorm:"type:varchar(100)"`
	LockTime     *time.Time
}
