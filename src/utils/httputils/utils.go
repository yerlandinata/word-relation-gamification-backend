package httputils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"

	jwt "github.com/dgrijalva/jwt-go"
)

type HTTPError struct {
	ErrorMessage string `json:"error_message"`
}

type ContextKey string

func GetPlayerIDFromJWT(r *http.Request) int64 {
	var ctxKey ContextKey = "player_id"
	return r.Context().Value(ctxKey).(int64)
}

func Authenticate(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		authHeaderSplit := strings.Split(authHeader, " ")
		tokenString := authHeaderSplit[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(config.GetAppConfig().Secret), nil
		})

		if err != nil {
			ErrorResponseJSON(w, http.StatusInternalServerError, err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			id := claims["player_id"]
			var contextKey ContextKey = "player_id"
			playerID, err := strconv.ParseInt(strings.TrimSpace(id.(string)), 10, 64)

			if err != nil {
				ErrorResponseJSON(w, http.StatusInternalServerError, err)
				return
			}

			ctx := context.WithValue(r.Context(), contextKey, playerID)
			handler(w, r.WithContext(ctx))
		} else {
			ErrorResponseJSON(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		}

	}
}

func HandlePost(pattern string, handler http.HandlerFunc) {
	handleSpecificMethod(http.MethodPost)(pattern, handler)
}

func HandleGet(pattern string, handler http.HandlerFunc) {
	handleSpecificMethod(http.MethodGet)(pattern, handler)
}

func handleSpecificMethod(method string) func(string, http.HandlerFunc) {
	return func(pattern string, handler http.HandlerFunc) {
		http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodOptions {
				handleOptions(w, method)
				return
			} else if r.Method == method {
				handler(w, r)
				return
			}
			ErrorResponseJSON(w, http.StatusMethodNotAllowed, errors.New("Method not allowed"))
		})
	}
}

func handleOptions(w http.ResponseWriter, allowedMethod string) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", allowedMethod)
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.WriteHeader(http.StatusOK)
}

func ResponseJSON(w http.ResponseWriter, statusCode int, body interface{}) {

	js, err := json.Marshal(body)
	if err != nil {
		ErrorResponseJSON(w, http.StatusInternalServerError, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	w.Write(js)

}

func ErrorResponseJSON(w http.ResponseWriter, statusCode int, err error) {
	js, err := json.Marshal(HTTPError{ErrorMessage: err.Error()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.WriteHeader(statusCode)
	w.Write(js)
}
