package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/labadiejoaco/gin-gorm-api/database"
	"github.com/labadiejoaco/gin-gorm-api/models"
	"gorm.io/gorm"
)

type body struct {
	Title string `json:"title" validate:"required,min=2,max=20"`
	Body  string `json:"body" validate:"required,min=10,max=100"`
}

func GetPosts(c *gin.Context) {
	var posts []models.Post

	res := database.DB.Find(&posts)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Posts not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  posts,
		"posts": res.RowsAffected,
	})
}

func GetPostById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	res := database.DB.First(&post, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func CreatePost(c *gin.Context) {
	var body body

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	validate := validator.New()

	err := validate.Struct(body)
	if err != nil {
		// for _, err := range err.(validator.ValidationErrors) {
		// 	fmt.Println(err.Error())
		// }

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}

	res := database.DB.Create(&post)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not create post",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")

	var body body

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	validate := validator.New()

	err := validate.Struct(body)
	if err != nil {
		// for _, err := range err.(validator.ValidationErrors) {
		// 	fmt.Println(err.Error())
		// }

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	var post models.Post

	res := database.DB.First(&post, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}

	res = database.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not update post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": post,
	})
}

func DeletePosts(c *gin.Context) {
	res := database.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Post{})
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not delete posts",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Posts deleted",
	})
}

func DeletePostById(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	res := database.DB.First(&post, id)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Post not found",
		})
		return
	}

	res = database.DB.Delete(&models.Post{}, id)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not delete post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted",
	})
}
