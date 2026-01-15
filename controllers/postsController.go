package controllers

import (
	"github.com/chirag0785/go-tut-api/dto"
	"github.com/chirag0785/go-tut-api/initializers"
	"github.com/chirag0785/go-tut-api/models"
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"net/http"
)

func PostsCreate(c *gin.Context) {

	//Get data off req body

	val, _ := c.Get("body")
	user_id, _ := c.Get("user_id")

	bodyData, _ := val.(*dto.PostCreateDTO)
	//create post

	post := models.Post{Title: bodyData.Title, Body: bodyData.Body, UserID: user_id.(uint)}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create post",
		})
		return
	}
    c.JSON(http.StatusOK, gin.H{
       "post": post,
    })
}

func PostsIndex(c *gin.Context) {

	//get the posts

	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	//respond with them
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {

	//Get id off url
	id := c.Param("id")

	var post models.Post

	result := initializers.DB.First(&post,id)
	if result.Error != nil {
		
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError,gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	
	//Get id off url

	user_id, _ := c.Get("user_id")

	id := c.Param("id")
	var body dto.PostUpdateDTO

	if err:= c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	var post models.Post

	result := initializers.DB.First(&post,id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError,gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	if post.UserID != user_id.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you are not allowed to update this post",
		})
		return
	}

	if body.Title != nil {
		post.Title= *body.Title
	}
	if body.Body != nil {
		post.Body= *body.Body
	}

	initializers.DB.Save(&post)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostsDelete(c * gin.Context) {

	//Get id off url

	user_id, _ := c.Get("user_id")

	id := c.Param("id")

	var post models.Post

	result := initializers.DB.First(&post,id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": result.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError,gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	
	if post.UserID != user_id.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "you are not allowed to delete this post",
		})
		return
	}

	results := initializers.DB.Delete(&models.Post{},id)

	if results.Error != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"error": "failed to delete the post",
		})
		return
	}

	if results.RowsAffected == 0 {
		c.JSON(http.StatusNotFound,gin.H{
			"error": "No post found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "post deleted successfully",
	})

}

func PostsUserPosts(c *gin.Context) {
	//get user id from context
	user_id, _ := c.Get("user_id")

	var posts []models.Post

	result := initializers.DB.Where(models.Post{UserID: user_id.(uint)}).Find(&posts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}