package controllers

import (
	"encoding/json"
	"libary/utility"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"` // Admin or Regular
}

// JWTToken represents the structure of JWT token
type JWTToken struct {
	Token string `json:"token"`
}

var secretKey = []byte("secret") // Change this with a stronger secret key in production

func Login(w http.ResponseWriter, r *http.Request) {

	log.Println("hello login")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Authenticate user (sample data)
	if user.Username == "admin" && user.Password == "adminpassword" {
		// Admin user
		tokenString, err := GenerateJWTToken("admin")
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}
		response := JWTToken{Token: tokenString}
		json.NewEncoder(w).Encode(response)
	} else if user.Username == "regular" && user.Password == "regularpassword" {
		// Regular user
		tokenString, err := GenerateJWTToken("regular")
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}
		response := JWTToken{Token: tokenString}
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}

}

func GenerateJWTToken(userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_type": userType,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GetUserType(r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)
	userType := claims["user_type"].(string)

	return userType
}

func Forbidden(w http.ResponseWriter, r *http.Request) {
	utility.RenderTemplate(w, r, "page_403", nil)
}
