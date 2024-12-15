package userinfo

import (
	"auth/verify"
	"auth/xsupabase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserInfo struct {
	UserId     string `json:"user_id"`
	CreatedAt  string `json:"created_at"`
	NationalId string `json:"national_id"`
	Email      string `json:"email"`
}

const (
	Table = "user_info"
)

// Create a new user's information row
func Create(c *gin.Context) {
	SupabaseClient := xsupabase.SupabaseClient
	var user UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetString("userID")
	user.UserId = userID
	if !verify.NationalID(user.NationalId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NationalID is not valid"})
		return
	}

	if !verify.Email(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is not valid"})
		return
	}

	User := map[string]string{
		"user_id":     user.UserId,
		"email":       user.Email,
		"national_id": user.NationalId,
	}

	log.Printf(">> User ID: %q\n", user.UserId)

	result, count, err := SupabaseClient.From(Table).
		Insert(User, false, "", "", "").
		Execute()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to add user info", "details": err.Error()})
		return
	}

	log.Printf("Create user_info: %v %v %v\n", result, count, err)
	c.JSON(http.StatusCreated, gin.H{"message": "User information created successfully", "data": result})
}

// Read all information for the authenticated user
func Read(c *gin.Context) {
	SupabaseClient := xsupabase.SupabaseClient

	userID := c.GetString("userID")
	result, count, err := SupabaseClient.From(Table).
		Select("*", "", false).
		Eq("user_id", userID).
		Execute()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read user info", "details": err.Error()})
		return
	}

	log.Printf("Read user_info: %v %v %v\n", result, count, err)
	c.JSON(http.StatusCreated, gin.H{"message": "User information read successfully", "data": result})
}

// Update an existing user's information
func Update(c *gin.Context) {
	SupabaseClient := xsupabase.SupabaseClient

	var user UserInfo
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetString("userID")
	user.UserId = userID
	if !verify.NationalID(user.NationalId) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "NationalID is not valid"})
		return
	}

	if !verify.Email(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is not valid"})
		return
	}

	User := map[string]string{
		"user_id":     user.UserId,
		"email":       user.Email,
		"national_id": user.NationalId,
	}

	result, count, err := SupabaseClient.From(Table).
		Update(User, "", "").
		Execute()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update user info", "details": err.Error()})
		return
	}

	log.Printf("Update user_info: %v %v %v\n", result, count, err)
	c.JSON(http.StatusCreated, gin.H{"message": "User information updated successfully", "data": result})
}


type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}


// login handles user login
func Login(c *gin.Context){
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":"invalid request payload"})
		return
	}

	authResponse , err := xsupabase.SupabaseClient.Auth.SignInWithEmailPassword(req.Email,req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login success",
		"token": authResponse.AccessToken,
	})
}