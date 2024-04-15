package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Define a secret key for signing JWT tokens
var jwtKey = []byte("your-secret-key")

// User represents a user object
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims represents JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Handler for generating JWT token upon successful login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if user credentials are valid (dummy check for demo)
	if user.Username == "user" && user.Password == "password" {
		expirationTime := time.Now().Add(24 * time.Hour)
		claims := &Claims{
			Username: user.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Login successful")
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}

// Middleware function to authenticate requests using JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tokenString := cookie.Value
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to next handler
		next.ServeHTTP(w, r)
	})
}

// Protected route that requires authentication
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Protected route accessed successfully")
}

// func main() {
// 	// Register routes
// 	http.HandleFunc("/login", loginHandler)
// 	http.Handle("/protected", authMiddleware(http.HandlerFunc(protectedHandler)))

// 	// Start server
// 	fmt.Println("Server is listening on port 8080...")
// 	http.ListenAndServe(":8080", nil)
// }
