package usecase

import (
	"errors"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
	"github.com/yerlandinata/word-relation-gamification-backend/src/domain"
)

const (
	WrongUserIDOrPassword int = 3
	UserDoesNotExists     int = 2
	LoginOK               int = 1
)

func Login(userID int64, password int64) (*domain.Player, int, error) {
	player, err := domain.GetPlayerByID(userID)

	if err != nil {
		return nil, 0, err
	}

	if player == nil {
		return nil, UserDoesNotExists, nil
	}

	if player.Password != password {
		return nil, WrongUserIDOrPassword, nil
	}

	return player, LoginOK, nil
}

func Register(player *domain.Player) error {
	return domain.AddPlayer(player)
}

func UpdatePlayerID(player *domain.Player, newID int64) error {
	return domain.UpdatePlayerID(player, newID)
}

func ResetPlayerScoreAndTime(playerID int64) error {
	// first, please check if the player can play again
	// player can play again if "we think" that the potential score could be higher

	annotationCriteria := domain.AnnotationCriteria{
		NotAnnotatedByPlayerID: playerID,
		MaxCount:               config.GetAppConfig().TargetAnnotationCountPerWordPair,
	}

	// count the non gold standard
	annotationCriteria.IsGoldStandard = false
	wordPairs, err := domain.GetWordPairByAnnotationCriteria(annotationCriteria, 100)
	if err != nil {
		return err
	}

	if len(wordPairs) < 100 {
		return errors.New("cannot play anymore")
	}

	// count the gold standard
	annotationCriteria.IsGoldStandard = true
	wordPairs, err = domain.GetWordPairByAnnotationCriteria(annotationCriteria, 25)
	if err != nil {
		return err
	}

	if len(wordPairs) < 25 {
		return errors.New("cannot play anymore")
	}

	return domain.ResetPlayerScoreAndTime(playerID)
}

func GetRankingsNearPlayer(player *domain.Player) ([]domain.Player, error) {
	player, err := domain.GetPlayerByID(player.ID)
	if err != nil {
		return nil, err
	}

	players, err := domain.GetPlayerRankingsRange(player.Rank-5, player.Rank)
	if err != nil {
		return nil, err
	}

	return players, nil
}
