package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jblaski/go-jwt/database"
	"github.com/jblaski/go-jwt/models"
	"gorm.io/gorm"
)

// GetProfile returns user data
func GetProfile(c *gin.Context) {
	var user models.User

	email, _ := c.Get("email")

	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get user profile",
		})
		c.Abort()
		return
	}

	user.Password = ""

	c.JSON(200, user)
}

func GetSecret(c *gin.Context) {
	var user models.User

	email, _ := c.Get("email")

	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get user profile",
		})
		c.Abort()
		return
	}

	if user.Secret == "" {
		c.JSON(200, gin.H{
			"msg": "user has no secret set!",
		})
		c.Abort()
		return
	}

	c.JSON(200, user.Secret)
}

type SecretPayload struct {
	Secret string `json:"secret"`
}

// New functionality
func PutSecret(c *gin.Context) {
	//TODO: what should a PUT reply with, if anything?
	var user models.User
	email, _ := c.Get("email") // get the email of the user

	var secretPayload SecretPayload // get the secret from the supplied message body
	err := c.ShouldBindJSON(&secretPayload)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user) // get user from DB using email

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get user profile",
		})
		c.Abort()
		return
	}

	user.UpdateUserSecret(secretPayload.Secret) // update their secret
}
