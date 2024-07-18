package models

import (
	"github.com/lib/pq"
)

type Word struct {
	ID           uint           `gorm:"primarykey"`
	English      string         `gorm:"unique"`
	Hashtags     pq.StringArray `gorm:"type:text[]"`
	Learned      bool           `gorm:"default:false"`
	Translations []Translation  `gorm:"foreignKey:WordRefer"`
	CreatedAt    string
}
