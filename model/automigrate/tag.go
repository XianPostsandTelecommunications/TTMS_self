package automigrate

import "gorm.io/gorm"

type TagType []string

// Tag  电影标签
type Tag struct {
	gorm.Model
	MovieID uint
	Movie   Movie
	Tags    TagType
}
