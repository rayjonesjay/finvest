package routes

import (
	"api/database"
	"api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var auth models.Auth

	if err := c.ShouldBind(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(auth.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to signup"})
		return
	}

	user := map[string]interface{}{
		"username": auth.Username,
		"email":    auth.Email,
		"password": string(hashedPassword),
		"phone":    auth.Phone,
	}

	// insert user into supabase
	if err := database.InsertUser(user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "signup success"})
}

func Login(c *gin.Context) {
	var auth models.Auth

	if err := c.ShouldBind(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fetcg yser from supabase by email
	user, err := database.GetUserByEmail(auth.Email)
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user["password"].(string)), []byte(auth.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid crentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "log successful"})
}
