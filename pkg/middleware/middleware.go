package middleware

import (
	"context"
	"net/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
	"github.com/aditya-411/mvc_assignment/pkg/types"
	"strings"
)

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noAuthPaths := []string{"/login", "/register", "/"}
		if stringInSlice(r.URL.Path, noAuthPaths){
			cookie, err := r.Cookie("login_token")
			if err != nil{
				next.ServeHTTP(w, r)
				return
			}


			err = verifyToken(cookie.Value)
			if err == nil {
				claims, pass := extractClaims(cookie.Value)
				if !pass {
					next.ServeHTTP(w, r)
					return
				}

				username := claims["username"].(string)
				isadmin := claims["isadmin"].(bool)
				r = r.WithContext(context.WithValue(r.Context(), types.Key("username"), username))
				r = r.WithContext(context.WithValue(r.Context(), types.Key("isadmin"), isadmin))
				http.Redirect(w, r, "/user", http.StatusFound)
				return
			}

			next.ServeHTTP(w, r)
			return
		} else {
			cookie, err := r.Cookie("login_token")
			if err != nil{
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}


			err = verifyToken(cookie.Value)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}


			claims, pass := extractClaims(cookie.Value)
			if !pass {
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}


			username := claims["username"].(string)
			isadmin := claims["isadmin"].(bool)
			r = r.WithContext(context.WithValue(r.Context(), types.Key("username"), username))
			r = r.WithContext(context.WithValue(r.Context(), types.Key("isadmin"), isadmin))
			route := r.URL.Path
			if strings.HasPrefix(route, "/admin") && !isadmin {
				http.Redirect(w, r, "/user", http.StatusFound)
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}


func verifyToken(tokenString string) error {
	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		return err
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Get the JWT secret from environment variable
		secret := os.Getenv("jwt_secret")
		if secret == "" {
			return nil, fmt.Errorf("JWT secret not found")
		}
		return []byte(secret), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func extractClaims(tokenStr string) (jwt.MapClaims, bool) {
	err := godotenv.Load(filepath.Join("./", ".env"))
	if err != nil {
		return nil, false
	}
	hmacSecretString := os.Getenv("jwt_secret")
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		 // check token signing method etc
		 return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		fmt.Printf("Invalid JWT Token")
		return nil, false
	}
}
