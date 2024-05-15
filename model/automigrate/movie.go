package automigrate

import "gorm.io/gorm"

type ActorString []string

// Movie  电影结构体
type Movie struct {
	gorm.Model
	Name       string      `gorm:"type:varchar(255);not null"`  // 电影名
	Area       string      `gorm:"type:varchar(255);not null"`  // 电影地区
	Actors     ActorString `gorm:"type:varchar(255);not null"`  // 演员表
	Avatar     string      `gorm:"type:varchar(255);not null"`  // 电影封面链接
	Content    string      `gorm:"type:varchar(1000);not null"` // 电影简介
	Duration   int64       `gorm:"type:int;not null"`           // 电影时长
	ShowTime   string      `gorm:"type:varchar(1000);not null"` // 上映时间
	Director   string      `gorm:"type:varchar(255);not null"`  // 导演
	Score      float32     `gorm:"type:float;not null"`         // 评分
	BoxOffice  float32     `gorm:"type:float;not null"`         // 票房输入
	VisitCount int         `gorm:"type:int;not null"`           // 观看次数
}
