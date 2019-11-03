package httphandler

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
	"github.com/yerlandinata/word-relation-gamification-backend/src/domain"
	"github.com/yerlandinata/word-relation-gamification-backend/src/usecase"
	"github.com/yerlandinata/word-relation-gamification-backend/src/utils/httputils"

	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	UserID   int64 `json:"id"`
	Password int64 `json:"birth_date"`
}

type LoginResponse struct {
	LoginStatus int            `json:"login_status"`
	Token       string         `json:"token"`
	Player      *domain.Player `json:"user,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginRequest)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusBadRequest, err)
		return
	}

	player, loginStatus, err := usecase.Login(loginRequest.UserID, loginRequest.Password)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	var loginResponse LoginResponse

	if player != nil {
		token, err := buildToken(player)

		if err != nil {
			httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
			return
		}

		loginResponse = LoginResponse{LoginStatus: loginStatus, Player: player, Token: token}

	} else {
		loginResponse = LoginResponse{LoginStatus: loginStatus}
	}

	httputils.ResponseJSON(w, http.StatusOK, loginResponse)

}

func Register(w http.ResponseWriter, r *http.Request) {
	var player domain.Player
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&player)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusBadRequest, err)
		return
	}

	err = usecase.Register(&player)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	token, err := buildToken(&player)

	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	loginResponse := LoginResponse{LoginStatus: usecase.LoginOK, Token: token}
	httputils.ResponseJSON(w, http.StatusOK, loginResponse)
}

func buildToken(player *domain.Player) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"player_id": fmt.Sprintf("%d", player.ID) + " ",
		"password":  fmt.Sprintf("%d", player.Password) + " ",
	})

	tokenStr, err := token.SignedString([]byte(config.GetAppConfig().Secret))
	return tokenStr, err
}
