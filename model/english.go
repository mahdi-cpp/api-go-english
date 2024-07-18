package model

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
	"time"
)

type Word struct {
	ID           uint           `gorm:"primarykey"`
	English      string         `gorm:"unique"`
	Hashtags     pq.StringArray `gorm:"type:text[]"`
	Learned      bool           `gorm:"default:false"`
	Translations []Translation  `gorm:"foreignKey:WordRefer"`
	CreatedAt    time.Time
}

type Translation struct {
	ID        uint           `gorm:"primarykey"`
	Persians  pq.StringArray `gorm:"type:text[]"`
	Type      string         `gorm:"default:null"`
	WordRefer uint
	CreatedAt time.Time
}

type Category struct {
	gorm.Model
	Hashtag string `gorm:"unique"`
	Learn   string `gorm:"default:all"`
	Order   string `gorm:"default:ASC"`
	Type    string `gorm:"default:null"`
	Page    uint64 `gorm:"default:0"`
}

type Filters struct {
	Learn   string
	Type    string
	Hashtag string
	Order   string
}

type EnglishEntity struct {
	//Category Category
	Words []Word
	Count int64
}

type User2 struct {
	ID       uint   `gorm:"primarykey"`
	Username string `gorm:"unique"`
	Email    string
	Phone    string
	Persians pq.StringArray `gorm:"type:text[][]"`
}

type User struct {
	gorm.Model
	Username   string `gorm:"unique"`
	Email      string
	Phone      string
	Avatar     string
	FullName   string
	Biography  string
	IsVerified bool
	Learn      string `gorm:"default:all"`
	Order      string `gorm:"default:ASC"`
	Type       string `gorm:"default:null"`
	Hashtag    string `gorm:"default:null"`
	Page       uint64 `gorm:"default:0"`
}
