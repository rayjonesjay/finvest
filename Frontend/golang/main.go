package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", HomePageHandler)
	http.HandleFunc("/about", AboutPageHandler)
	http.HandleFunc("/register", RegisterPageHandler)
	http.HandleFunc("/sign-in", SignInPageHandler)
	http.HandleFunc("/login", LoginPageHandler)
	http.HandleFunc("/borrow", BorrowPageHandler)
	http.HandleFunc("/lend", LendPageHandler)
	http.HandleFunc("/profile", UserProfilePageHandler)


	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/home.html")
}

func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/about.html")
}

func RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/register.html")
}

func SignInPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/sign_in.html")
}

func LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/login.html")
}

func BorrowPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/loan_request.html")
}

func LendPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/lend.html")
}

func UserProfilePageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/user_profile.html")
}

// Dummy user data for demonstration
var validUser = map[string]string{
	"email":    "user@example.com",
	"password": "password123",
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse incoming JSON request body
		var userData map[string]string
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Check if the login details match
		if userData["email"] == validUser["email"] && userData["password"] == validUser["password"] {
			// Successful login
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Login successful!"})
		} else {
			// Failed login
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"message": "Invalid email or password"})
		}
	}
}
