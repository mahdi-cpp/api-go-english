package main

import (
	"github.com/mahdi-cpp/api-go-english/tree/api_v1/repository"
)

func main() {

	repository.Init()
	//repository.InitUser()
	//repository.InitCategory()
	//repository.CreatPasswords()

	Run()
}

func upgradeDatabase() {

	//englishes := repository.GetEnglishMultiPersian()
	//for i := 0; i < len(englishes); i++ {
	//	parts := strings.Split(englishes[i].Persian1, "-")
	//	fmt.Println(parts[0])
	//	repository.SetEnglishUpdatePersians(englishes[i].ID, parts[0])
	//}

	//englishes := repository.GetEnglishAllRows()
	//for i := 0; i < len(englishes); i++ {
	//	var tags []string
	//	if englishes[i].University == true {
	//		tags = append(tags, "University")
	//	}
	//	if englishes[i].AEF == true {
	//		tags = append(tags, "American English File")
	//	}
	//	if englishes[i].Oxford == true {
	//		tags = append(tags, "Oxford")
	//	}
	//	if englishes[i].Finance == true {
	//		tags = append(tags, "Finance")
	//	}
	//	if englishes[i].Medium == true {
	//		tags = append(tags, "Medium")
	//	}
	//	if englishes[i].Programing == true {
	//		tags = append(tags, "Programing")
	//	}
	//	if englishes[i].Word504 == true {
	//		tags = append(tags, "Word504")
	//	}
	//	//fmt.Println(tags)
	//	err := repository.SetEnglishUpdateTags(englishes[i].ID, tags)
	//	if err != nil {
	//		return
	//	}
	//}

	//englishes := repository.GetEnglishAllRows()
	//for i := 0; i < len(englishes); i++ {
	//	var persians []string
	//	if  len(englishes[i].Persian1) > 0 {
	//		persians = append(persians, englishes[i].Persian1)
	//	}
	//	if  len(englishes[i].Persian2) > 0 {
	//		persians = append(persians, englishes[i].Persian2)
	//	}
	//	if  len(englishes[i].Persian3) > 0 {
	//		persians = append(persians, englishes[i].Persian3)
	//	}
	//	if  len(englishes[i].Persian4) > 0 {
	//		persians = append(persians, englishes[i].Persian4)
	//	}
	//	if  len(englishes[i].Persian5) > 0 {
	//		persians = append(persians, englishes[i].Persian5)
	//	}
	//
	//	//fmt.Println(persians)
	//	err := repository.SetEnglishUpdatePersians(englishes[i].ID, persians)
	//	if err != nil {
	//		return
	//	}
	//}

}
