package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Hashtag string `gorm:"unique"`
	Learn   string `gorm:"default:all"`
	Order   string `gorm:"default:ASC"`
	Type    string `gorm:"default:null"`
	Page    uint64 `gorm:"default:0"`
}
