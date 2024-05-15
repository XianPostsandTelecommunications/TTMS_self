package dao

import (
	"gorm.io/gorm"
	redis "mognolia/internal/dao/redis/query"
)

type group struct {
	DB    *gorm.DB
	Redis *redis.Queries
}

var Group = new(group)
