package usecase

import (
	"errors"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
	"github.com/yerlandinata/word-relation-gamification-backend/src/domain"
)

const (
	UserDoesNotExists int = 2
	LoginOK           int = 1
)

type PlayerGoldStandardAgreements struct {
	OverallAgree      int
	OverallDisagree   int
	HyponymyAgree     int
	HyponymyDisagree  int
	UnrelatedAgree    int
	UnrelatedDisagree int
}

func Login(userID int64) (*domain.Player, int, error) {
	player, err := domain.GetPlayerByID(userID)

	if err != nil {
		return nil, 0, err
	}

	if player == nil {
		return nil, UserDoesNotExists, nil
	}

	return player, LoginOK, nil
}

func Register(player *domain.Player) error {
	return domain.AddPlayer(player)
}

func UpdatePlayerIDAndName(player *domain.Player, newID int64, newName string) error {
	return domain.UpdatePlayerIDAndName(player, newID, newName)
}

func RecordPlayerOnboardingTime(player *domain.Player, onboardingTimeMS int) error {
	return domain.SetPlayerOnboardingTime(player, onboardingTimeMS)
}

func PlayerLevelUp(playerID int64) error {

	annotationCriteria := domain.AnnotationCriteria{
		NotAnnotatedByPlayerID: playerID,
		MaxCount:               config.GetAppConfig().TargetAnnotationCountPerWordPair,
	}

	// count the non gold standard
	annotationCriteria.IsGoldStandard = false
	wordPairs, err := domain.GetWordPairByAnnotationCriteria(annotationCriteria, 1)
	if err != nil {
		return err
	}

	if len(wordPairs) < 1 {
		return errors.New("cannot play anymore")
	}

	// count the gold standard
	annotationCriteria.IsGoldStandard = true
	annotationCriteria.MaxCount = config.GetAppConfig().TargetAnnotationCountPerGoldStandard
	wordPairs, err = domain.GetWordPairByAnnotationCriteria(annotationCriteria, 1)
	if err != nil {
		return err
	}

	if len(wordPairs) < 1 {
		return errors.New("cannot play anymore")
	}

	return domain.IncrementPlayerLevelAndResetPlayerTime(playerID)
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
