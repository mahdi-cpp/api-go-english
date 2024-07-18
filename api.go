package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mahdi-cpp/api-go-english/models"
	"github.com/mahdi-cpp/api-go-english/repository"
	"strconv"
	"strings"
)

func addEnglishRoutes(rg *gin.RouterGroup) {

	router := rg.Group("/duolingo")

	router.POST("/addWord", func(c *gin.Context) {
		word := models.Word{}
		s := c.Query("data")
		//fmt.Println(s)

		err := PrettyJson(s, &word)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, err.Error())
			return
		}

		word.English = strings.ToLower(word.English)
		fmt.Println(word)
		fmt.Println("word.Hashtags: ", word.Hashtags)

		err = repository.AddWord(word)
		if err != nil {
			result := "" + err.Error()
			c.JSON(404, result)
		} else {
			c.JSON(200, "success")
		}
	})

	router.POST("/EditWord", func(c *gin.Context) {
		word := models.Word{}
		s := c.Query("word")

		err := PrettyJson(s, &word)
		if err != nil {
			fmt.Println(err)
			c.JSON(400, err.Error())
			return
		}
		err = repository.EditWord(word)
		if err != nil {
			c.JSON(404, err.Error())
		} else {
			c.JSON(200, "Successful edit '"+word.English+"'")
		}
	})

	//router.POST("/CategoryEdit", func(c *gin.Context) {
	//	category := models.Category{}
	//	err := PrettyJson(c.Query("category"), &category)
	//	if err != nil {
	//		fmt.Println(err)
	//		c.JSON(400, err.Error())
	//		return
	//	}
	//
	//	err = repository.EditCategory(category)
	//	if err != nil {
	//		result := "" + err.Error()
	//		fmt.Println("EditUser Error: ", result)
	//		c.JSON(400, result)
	//	} else {
	//		fmt.Println("category edit success")
	//		c.JSON(200, "success")
	//	}
	//})

	router.POST("/UpdateLearned", func(c *gin.Context) {
		id := c.Query("id")
		checked := c.Query("checked")
		learned, _ := strconv.ParseBool(checked)
		err := repository.EditLearned(id, learned)
		if err != nil {
			c.JSON(404, err.Error())
		} else {
			c.JSON(200, "Successful update")
		}
	})

	router.POST("/DeleteById", func(c *gin.Context) {
		word := models.Word{}
		s := c.Query("data")
		fmt.Println(s)
		err := json.Unmarshal([]byte(s), &word)
		if err != nil {
			fmt.Println(err)
			result := "Unmarshal" + err.Error()
			c.JSON(200, result)
			return
		}
		err = repository.DeleteByWord(word)
		if err != nil {
			return
		}
	})

	router.POST("/editLearn", func(c *gin.Context) {
		hashtag := c.Query("hashtag")
		learn := c.Query("learn")
		err := repository.EditLearn(hashtag, learn)

		if err != nil {
			fmt.Println("editLearn Error: ", err.Error())
			c.JSON(404, err.Error())
			return
		}
		fmt.Println("editLearn success")
		c.JSON(200, "successful")
	})

	router.POST("/editOrder", func(c *gin.Context) {
		hashtag := c.Query("hashtag")
		order := c.Query("order")
		err := repository.EditOrder(hashtag, order)
		if err != nil {
			c.JSON(404, err)
			return
		}
		c.JSON(200, "successful")
	})

	router.POST("/editType", func(c *gin.Context) {
		hashtag := c.Query("hashtag")
		kind := c.Query("type")
		err := repository.EditType(hashtag, kind)
		if err != nil {
			c.JSON(404, err)
			return
		}
		c.JSON(200, "successful")
	})

	router.POST("/editPage", func(c *gin.Context) {
		hashtag := c.Query("hashtag")
		page := c.Query("page")
		err := repository.EditPage(hashtag, page)
		if err != nil {
			c.JSON(404, err)
			return
		}
		c.JSON(200, "successful")
	})

	router.GET("/getCategory", func(c *gin.Context) {
		hashtag := c.Query("hashtag")
		category, err := repository.GetCategory(hashtag)
		if err != nil {
			c.JSON(404, err.Error())
			return
		}
		c.JSON(200, category)
	})

	router.GET("/GetWords", func(c *gin.Context) {
		hashtag := c.Query("hashtag")
		search := c.Query("search")
		words, err := repository.GetWords(hashtag, search)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(417, err.Error())
			return
		}
		c.JSON(200, words)
	})

	router.GET("/GetById", func(c *gin.Context) {
		words, err := repository.GetById(c.Query("id"))
		if err != nil {
			c.JSON(404, err.Error())
		}
		c.JSON(200, words)
	})

	router.GET("/GetIsWordAvailable", func(c *gin.Context) {
		word := c.Query("word")
		var count = repository.GetIsWordAvailable(word)
		fmt.Println(count)

		if count == 0 {
			c.JSON(200, "Not")
		} else {
			c.JSON(200, "Available")
		}
	})

	router.GET("/getTest", func(c *gin.Context) {
		word := c.Query("word")
		var count = repository.GetIsWordAvailable(word)
		fmt.Println(count)

		if count == 0 {
			c.JSON(200, "Not")
		} else {
			c.JSON(200, "Available")
		}
	})
}

func PrettyJson(param string, T interface{}) error {
	err := json.Unmarshal([]byte(param), T)

	if err != nil {
		fmt.Println(err)
		return err
	}
	b, err := json.MarshalIndent(T, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

	return nil
}
