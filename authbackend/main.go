package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"auth/routes/userinfo"
	"auth/xsupabase"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/supabase-community/supabase-go"
)

// Note struct for CRUD operations
type Note struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	// Initialize Supabase client
	supabaseURL := os.Getenv("PUBLIC_SUPABASE_URL")
	supabaseKey := os.Getenv("PUBLIC_SUPABASE_ANON_KEY")

	xsupabase.SupabaseClient, err = supabase.NewClient(supabaseURL, supabaseKey, &supabase.ClientOptions{})
	if err != nil {
		log.Fatal("cannot initialize supabase client", err)
	}
	log.Println("[OK] supabase client initialized")
	// data, count, err := supabaseClient.From("countries").Select("*", "exact", false).Execute()

	r := gin.Default()
	if os.Getenv("DEVELOPMENT_MODE") == "TRUE" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// static files
	r.Static("/static", "../src/templates/")

	// Load HTML templates
	r.LoadHTMLGlob("../src/templates/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", nil)
	})

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	r.POST("/signup", func(c *gin.Context) {
		email := c.PostForm("email")
		password := c.PostForm("password")

		if email == "" || password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"email": "email and password required"})
			return
		}

		// TODO: add supabase auth here :)

		c.JSON(http.StatusOK, gin.H{"message": "signup success"})
	})

	// Routes
	r.Use(AuthMiddleware())
	{
		// Register User Information routes
		routeGroup := r.Group("/auth")
		{
			routeGroup.POST("/signup", userinfo.Create)
			routeGroup.POST("/login", userinfo.Login)
			routeGroup.GET("/read", userinfo.Read)
			routeGroup.GET("/update", userinfo.Update)
		}
	}

	// r.POST("/notes", createNote)
	// r.GET("/notes", getNotes)
	// r.PUT("/notes/:id", updateNote)
	// r.DELETE("/notes/:id", deleteNote)

	log.Fatal(r.Run(":8080"))
}

// AuthMiddleware Middleware to validate JWT and extract user ID
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		userClient := xsupabase.SupabaseClient.Auth.WithToken(token)
		user, err := userClient.GetUser()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		log.Printf("User ID: %v -> Token User: %v\n", user.User.ID, user)

		// Pass user ID to the context
		c.Set("userID", user.User.ID)
		c.Next()
	}
}
