package controllers

import (
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

func PutSecret(c *gin.Context) {
	//TODO: how to update value? What should a PUT reply with>
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

	user.Secret = "test"
	user.UpdateUserSecret()

}
