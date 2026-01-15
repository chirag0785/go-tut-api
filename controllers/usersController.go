package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/chirag0785/go-tut-api/dto"
	"github.com/chirag0785/go-tut-api/initializers"
	"github.com/chirag0785/go-tut-api/models"
	"github.com/chirag0785/go-tut-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func UsersCreate(c *gin.Context) {
	var body dto.UserCreateDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	//check if user with email already exists
	var existingUser models.User
	initializers.DB.Where(&models.User{Email: body.Email}).First(&existingUser)

	if existingUser.ID != 0 {		//user already exists
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user with this email already exists",
		})
		return
	}

	//hash password
	hashedPassword, err := utils.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to hash password",
		})
		return
	}
	//create user
	user := models.User{Name: body.Name, Email: body.Email, Password: hashedPassword}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UsersLogin(c *gin.Context) {
	//check if user with email exists
	var body dto.UserLoginDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	var user models.User
	result := initializers.DB.Where(&models.User{Email: body.Email}).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid email or password",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to fetch user",
		})
		return
	}

	//now check if password is correct
	isPasswordValid := utils.CheckPasswordHash(body.Password, user.Password)
	//if not correct -> return error

	if !isPasswordValid {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email or password",
		})
		return
	}

	//if correct -> generate jwt token and send it back
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email": user.Email,
		"exp":  time.Now().Add(time.Hour * 24).Unix(), //token valid for 7 days
	}
	token, err :=utils.GenerateJWTToken(os.Getenv("JWT_SECRET"),claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"message": "login successful",
	})
}
