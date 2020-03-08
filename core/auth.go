package core

import (
	"net/http"
	u "github.com/elcompadre/elcompadre/copro/utils"
	"strings"
	"github.com/elcompadre/elcompadre/copro/models"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"context"
	"fmt"
)

var JwtAuthentication = func(next http.Handler) http.Handler {
	return http.HandleFunc(func(writer http.ResponseWriter, request *http.Request){

		notAuth := []string{"/api/user/new", "/api/user/login"}
		requestPath := request.URL.Path //Current request path

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(writer, request)
				return
			}
		}

		response := make(map[string] interface{})
		tokenHeader := request.Header.Get("Authorization")

		if tokenHeader == "" {
			response = u.Message(false, "Missing auth token")
			writer.WriteHeader(http.StatusForbidden)
			u.Response(writer, response)
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response = u.Message(false, "Invalid/Malformed auth token")
			writer.WriteHeader(http.StatusForbidden)
			u.Response(writer, response)
			return
		}

		tokenPart := splitted[1]
		tk := &models.Token{}

		token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error){
			return []byte(os.Getenv("TOKEN_PASSWORD")), nil
		})

		if err != nil {
			response = u.Message(false, "Malformed authentication token")
			writer.WriteHeader(http.StatusForbidden)
			u.Response(writer, response)
			return
		}

		if !token.Valid {
			response = u.Message(false, "Token is not valid.")
			writer.WriteHeader(http.StatusForbidden)
			u.Response(writer, response)
			return
		}

		fmt.Sprintf("User %", tk.UserId)
		ctx := context.WithValue(request.Context(), tk.UserId , tk.UserId)
		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
	})
}