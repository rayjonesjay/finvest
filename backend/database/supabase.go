package database

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/supabase-community/postgrest-go"
)

var SupabaseClient *postgrest.Client

func InitSupabase(apiKey, supabaseURL string) {
	if apiKey == "" || supabaseURL == "" {
		log.Fatal("Supabase API key or URL is missing!")
	}

	SupabaseClient = postgrest.NewClient(supabaseURL, "public", nil)
	SupabaseClient.SetApiKey(apiKey)
	//SupabaseClient.SetAuthToken(apiKey)
	log.Println("Supabase client initialized successfully.")
}

func GetUserByEmail(email string) (map[string]interface{}, error) {
	response, err, er := SupabaseClient.From("users").Select("*", "", false).Eq("email", email).Single().Execute()
	if er != nil {
		log.Println("Error fetching user by email:", err, er)
		return nil, fmt.Errorf("failed to fetch user: %v, %v", err, er)
	}

	var user map[string]interface{}
	if parseErr := json.Unmarshal(response, &user); parseErr != nil {
		log.Println("Error parsing response:", parseErr)
		return nil, parseErr
	}

	return user, nil
}

func InsertUser(user map[string]interface{}) error {
	if user == nil {
		return fmt.Errorf("user data cannot be nil")
	}

	response, _, err := SupabaseClient.From("users").Insert(user, true, "", "", "").Execute()
	if err != nil {
		log.Println("Error inserting user:", err)
		return err
	}

	// Optionally, confirm insertion success from the response
	log.Println("User inserted successfully:", string(response))
	return nil
}
