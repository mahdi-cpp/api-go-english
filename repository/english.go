package repository

import (
	"fmt"
	"github.com/mahdi-cpp/api-go-english/tree/api_v1/config"
	"github.com/mahdi-cpp/api-go-english/tree/api_v1/model"
	"gorm.io/gorm"
	"strings"
	"time"
)

func InitUser() {
	db.Create(&model.User{Username: "mahdiabdolmaleki", Email: "mahdi.cpp@gmail.com", Phone: "09355512619", Avatar: "2018-10-23_13-55-58_UTC_profile_pic.jpg", Biography: "go lang programmer"})
}

var db *gorm.DB

func Init() {
	db = config.DB
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Word{})
	db.AutoMigrate(&model.Translation{})
	db.AutoMigrate(&model.Category{})
}

func InitCategory() {
	db.Create(&model.Category{Hashtag: "All"})
	db.Create(&model.Category{Hashtag: "University"})
	db.Create(&model.Category{Hashtag: "American English File"})
	db.Create(&model.Category{Hashtag: "Oxford"})
	db.Create(&model.Category{Hashtag: "Google"})
	db.Create(&model.Category{Hashtag: "Youtube"})
	db.Create(&model.Category{Hashtag: "Medium"})
	db.Create(&model.Category{Hashtag: "Instagram"})
	db.Create(&model.Category{Hashtag: "Podcast"})
	db.Create(&model.Category{Hashtag: "Electronic"})
	db.Create(&model.Category{Hashtag: "Finance"})
	db.Create(&model.Category{Hashtag: "Programing"})
	db.Create(&model.Category{Hashtag: "Word504"})
	db.Create(&model.Category{Hashtag: "Movie"})
	db.Create(&model.Category{Hashtag: "WestWorld"})
}

func AddWord(word model.Word) error {
	err := db.Debug().Create(&word).Error
	return err
}

func EditWord(word model.Word) error {

	var data = map[string]interface{}{
		"English":  strings.ToLower(word.English),
		"Hashtags": word.Hashtags,
		"Learned":  word.Learned,
	}

	err := db.Debug().Where("id", word.ID).Model(model.Word{}).Updates(data).Error
	if err != nil {
		return err
	}

	fmt.Println("Translation 1: ", word.Translations)

	t := word.Translations[0]
	err = db.Where("id", t.ID).
		Updates(&model.Translation{Persians: t.Persians, Type: t.Type}).Error
	if err != nil {
		return err
	}

	//init the loc
	loc, _ := time.LoadLocation("Asia/Tehran")

	//set timezone,
	now := time.Now().In(loc)

	if len(word.Translations) > 1 {
		fmt.Println("Translation 1: ", t.Persians)
		t = word.Translations[1]
		var data = map[string]interface{}{
			"WordRefer": word.ID,
			"Type":      t.Type,
			"Persians":  t.Persians,
			"CreatedAt": now,
		}
		if t.ID == 0 { //Add New
			err = db.Debug().Model(model.Translation{}).Create(data).Error
			if err != nil {
				return err
			}
		} else { //Edit
			var updateData = map[string]interface{}{
				"Type":     t.Type,
				"Persians": t.Persians,
			}
			err = db.Debug().Model(model.Translation{}).Where("id", t.ID).Updates(updateData).Error
			if err != nil {
				return err
			}
		}
	}

	return err
}

func EditLearned(id string, checked bool) error {
	var update = map[string]interface{}{"learned": checked}
	return db.Debug().Model(model.Word{}).Where("id", id).Updates(update).Error
}

func DeleteByWord(word model.Word) error {
	err := db.Debug().Where("id", word.ID).Delete(&model.Word{}).Error
	return err
}

func EditLearn(hashtag string, learn string) error {
	var update = map[string]interface{}{
		"Learn": learn,
		"Page":  0,
	}
	err := db.Model(model.Category{}).Where("hashtag", hashtag).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}
func EditOrder(hashtag string, order string) error {
	var update = map[string]interface{}{
		"Order": order,
	}
	err := db.Model(model.Category{}).Where("hashtag", hashtag).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}
func EditType(hashtag string, kind string) error {
	var update = map[string]interface{}{
		"Type": kind,
		"Page": 0,
	}
	err := db.Model(model.Category{}).Where("hashtag", hashtag).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}
func EditPage(hashtag string, page string) error {

	if hashtag == "All" { // Always show first page for search result
		page = "0"
	}

	var update = map[string]interface{}{
		"Page": page,
	}
	err := db.Model(model.Category{}).Where("hashtag", hashtag).Updates(update).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCategory(hashtag string) (model.Category, error) {
	var category model.Category
	err := db.Debug().Where("hashtag", hashtag).Find(&category).Error
	if err != nil {
		return category, err
	}
	fmt.Println(category)
	return category, nil
}
func GetWords(hashtag string, search string) (model.EnglishEntity, error) {

	var entity model.EnglishEntity
	var category model.Category
	var words []model.Word
	var where = ""
	var count int64

	fmt.Println("search: " + search)

	err := db.Debug().Where("hashtag", hashtag).Find(&category).Error
	if err != nil {
		return model.EnglishEntity{}, err
	}

	println("Category:", category.Hashtag)

	if strings.Contains(category.Learn, "learned") {
		where += "english.words.learned = true "
	} else if strings.Contains(category.Learn, "tutorial") {
		where += "english.words.learned = false"
	}

	if len(search) > 0 {
		if len(where) > 0 {
			where += " AND english.words.english like '" + search + "%'"
		} else {
			where += "english.words.english like '" + search + "%'"
		}
	}

	if hashtag != "All" {
		if len(where) > 0 {
			where += " AND  '" + hashtag + "'=ANY(english.words.hashtags)"
		} else {
			where = "'" + hashtag + "'=ANY(english.words.hashtags)"
		}
	}

	//if category.Type != "" {
	//	if len(where) > 0 {
	//		where += " AND english.translations.type = '" + category.Type + "'"
	//	} else {
	//		where = "english.translations.type = '" + category.Type + "'"
	//	}
	//}

	db.Debug().
		Preload("Translations", func(db *gorm.DB) *gorm.DB { return db.Order("id ASC") }).
		Where(where).
		Offset(int(category.Page * 5)).
		Limit(5).
		Order("id " + category.Order).
		Find(&words)

	db.Debug().Model(&model.Word{}).Where(where).Count(&count)

	entity.Words = words
	entity.Count = count

	println("......Count:", count)

	return entity, nil
}

func GetById(id string) (model.Word, error) {
	var word model.Word
	err := db.Debug().Where("id", id).
		Preload("Translations", func(db *gorm.DB) *gorm.DB { return db.Order("id ASC") }).
		Find(&word).Error
	return word, err
}

func GetIsWordAvailable(english string) int64 {
	var word model.Word
	result := db.Debug().Where("english", english).Find(&word)
	return result.RowsAffected
}
