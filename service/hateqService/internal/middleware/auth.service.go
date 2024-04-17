package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Souras/hateqnew/service/hateqService/internal/db"
	"github.com/Souras/hateqnew/service/hateqService/internal/models"
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
	// w.Header().Set("Content-Type", "application/json") // Example header
	// w.Header().Set("Access-Control-Allow-Origin", "*") // Example header (adjust as needed)
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	adminUsers, err := db.GetAdminUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	// Create a map to store the presence of each username
	// usernameMap := make(map[string]bool)

	// Populate the map with usernames from the data
	for _, dbUser := range adminUsers {
		// usernameMap[user.AdminID] = true
		if dbUser.AdminID == user.Username && dbUser.Pwd == user.Password {
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

			response := models.ResponseData{Status: true, Data: dbUser}
			jsonData, err := json.Marshal(response)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error marshalling JSON response: %v", err)
				return
			}

			// fmt.Fprintf(w, "Login successful")
			w.Write(jsonData)
			return
		}
	}

	http.Error(w, "Invalid username or password", http.StatusUnauthorized)
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
		// middleware.CorsMiddleware(w ,r)
		AddResponseHeader(w)
		// Token is valid, proceed to next handler
		next.ServeHTTP(w, r)
	})
}

// Protected route that requires authentication
// func protectedHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Protected route accessed successfully")
// }

// func main() {
// 	// Register routes
// 	http.HandleFunc("/login", loginHandler)
// 	http.Handle("/protected", authMiddleware(http.HandlerFunc(protectedHandler)))

// 	// Start server
// 	fmt.Println("Server is listening on port 8080...")
// 	http.ListenAndServe(":8080", nil)
// }
