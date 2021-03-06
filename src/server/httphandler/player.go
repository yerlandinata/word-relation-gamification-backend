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
	UserID int64 `json:"id"`
}

type PlayerDataUpdateRequest struct {
	NewID   int64  `json:"new_id"`
	NewName string `json:"full_name"`
}

type OnboardingRecordRequest struct {
	OnboardingTimeMs int `json:"onboarding_time_ms"`
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

	player, loginStatus, err := usecase.Login(loginRequest.UserID)
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

	httpStatus := http.StatusOK
	if loginResponse.LoginStatus != usecase.LoginOK {
		httpStatus = http.StatusForbidden
	}

	httputils.ResponseJSON(w, httpStatus, loginResponse)

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

func UpdatePlayerIDAndName(w http.ResponseWriter, r *http.Request) {
	var playerDataUpdateRequest PlayerDataUpdateRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&playerDataUpdateRequest)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusBadRequest, err)
		return
	}

	newID := playerDataUpdateRequest.NewID
	name := playerDataUpdateRequest.NewName
	player := httputils.GetPlayerFromJWT(r)

	err = usecase.UpdatePlayerIDAndName(player, newID, name)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	player.ID = newID

	token, err := buildToken(player)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	loginResponse := LoginResponse{LoginStatus: usecase.LoginOK, Token: token}
	httputils.ResponseJSON(w, http.StatusOK, loginResponse)
}

func PlayerLevelUp(w http.ResponseWriter, r *http.Request) {
	player := httputils.GetPlayerFromJWT(r)

	err := usecase.PlayerLevelUp(player.ID)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	wordPair, err := usecase.GetClassificationWordPair(player.ID)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	httputils.ResponseJSON(w, http.StatusOK, wordPair)
}

func RecordPlayerOnboardingTime(w http.ResponseWriter, r *http.Request) {
	player := httputils.GetPlayerFromJWT(r)

	var onboardingRecordRequest OnboardingRecordRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&onboardingRecordRequest)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusBadRequest, err)
		return
	}

	err = usecase.RecordPlayerOnboardingTime(player, onboardingRecordRequest.OnboardingTimeMs)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
	}

	httputils.ResponseJSON(w, http.StatusCreated, nil)
}

func GetRankingsNearPlayer(w http.ResponseWriter, r *http.Request) {
	player := httputils.GetPlayerFromJWT(r)

	players, err := usecase.GetRankingsNearPlayer(player)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	httputils.ResponseJSON(w, http.StatusOK, players)
}

func buildToken(player *domain.Player) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"player_id": fmt.Sprintf("%d", player.ID) + " ",
	})

	tokenStr, err := token.SignedString([]byte(config.GetAppConfig().Secret))
	return tokenStr, err
}
