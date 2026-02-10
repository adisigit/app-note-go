package controllers

import (
	"app-note-go/dto"
	"app-note-go/initializer"
	"app-note-go/models"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// @Tags Users
// @Param body body dto.RegisterUserRequest true "Register payload"
// @Router /users/register [post]
func CreateUser(c *gin.Context) {
	var body dto.RegisterUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{Username: body.Username, Email: body.Email, Password: body.Password}

	result := initializer.DB.Create(&user)
	if result.Error != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.AbortWithStatusJSON(200, gin.H{
		"message": "User created successfully",
	})
}

// @Tags Users
// @Param body body dto.LoginUserRequest true "Login payload"
// @Router /users/login [post]
func LoginUser(c *gin.Context) {
	var body dto.LoginUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	user := models.User{Email: body.Email}
	result := initializer.DB.Where("email = ?", body.Email).First(&user)
	if result.Error != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(200, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}
