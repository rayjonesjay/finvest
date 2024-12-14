package main

import (
	"api/database"
	"api/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envFile, _ := godotenv.Read(".env")

	supabaseURL := envFile["SUPABASE_URL"]
	supabaseKEY := envFile["SUPABASE_API"]

	if supabaseURL == "" || supabaseKEY == "" {
		log.Println("supabase url or url not set")
		return
	}

	// initialize supabase client
	database.InitSupabase(supabaseKEY, supabaseURL)

	router := gin.Default()

	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/signup", routes.Signup)
		authRoutes.POST("/login", routes.Login)
	}

	router.Run("localhost:8080")
}
