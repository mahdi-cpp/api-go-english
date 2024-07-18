package models

import (
	"github.com/lib/pq"
)

type Translation struct {
	ID        uint           `gorm:"primarykey"`
	Persians  pq.StringArray `gorm:"type:text[]"`
	Type      string         `gorm:"default:null"`
	WordRefer uint
	CreatedAt string
}
