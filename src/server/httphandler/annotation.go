package httphandler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yerlandinata/word-relation-gamification-backend/src/domain"
	"github.com/yerlandinata/word-relation-gamification-backend/src/usecase"
	"github.com/yerlandinata/word-relation-gamification-backend/src/utils/httputils"
)

type AddAnnotationResponse struct {
	Player   *domain.Player   `json:"player"`
	WordPair *domain.WordPair `json:"next_word_pair,omitempty"`
}

func AddAnnotation(w http.ResponseWriter, r *http.Request) {
	var annotation domain.Annotation
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&annotation)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusBadRequest, err)
		return
	}

	// look, do not trust the player ID from the request body, JWT is the truth
	annotation.PlayerID = httputils.GetPlayerFromJWT(r).ID

	player, err := usecase.AddAnnotation(&annotation)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}

	// the game must carry on
	wordPair, err := usecase.GetClassificationWordPair(player.ID)
	if err != nil {
		httputils.ErrorResponseJSON(w, http.StatusInternalServerError, err)
		return
	}
	if wordPair == nil {
		httputils.ErrorResponseJSON(w, http.StatusForbidden, fmt.Errorf("No word pair for player with ID %d", player.ID))
		return
	}

	httputils.ResponseJSON(w, http.StatusCreated, AddAnnotationResponse{Player: player, WordPair: wordPair})
}
