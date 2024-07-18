package models

import "github.com/lib/pq"

type New_Word struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt string

	English  string         `gorm:"unique"`
	Hashtags pq.StringArray `gorm:"type:text[]"`
	Learned  bool           `gorm:"default:false"`

	Persian1 Persian
	Persian2 Persian
}

type Persian struct {
	Persians pq.StringArray `gorm:"type:text[]"`
	Type     string         `gorm:"default:null"`
}
