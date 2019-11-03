package httphandler

import (
	"fmt"
	"net/http"

	"github.com/yerlandinata/word-relation-gamification-backend/src/usecase"
	"github.com/yerlandinata/word-relation-gamification-backend/src/utils/httputils"
)

func GetClassificationWordPair(w http.ResponseWriter, r *http.Request) {
	playerID := httputils.GetPlayerIDFromJWT(r)

	wordPair, err := usecase.GetClassificationWordPair(playerID)

	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	if wordPair == nil {
		httputils.ErrorResponseJSON(w, http.StatusForbidden, fmt.Errorf("No word pair for player with ID %d", playerID))
		return
	}

	httputils.ResponseJSON(w, http.StatusOK, wordPair)

}

func GetGoldStandards(w http.ResponseWriter, r *http.Request) {
	goldStandards, err := usecase.GetGoldStandards()
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	httputils.ResponseJSON(w, http.StatusOK, goldStandards)
}
