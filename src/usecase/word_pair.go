package usecase

import (
	"math/rand"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
	"github.com/yerlandinata/word-relation-gamification-backend/src/domain"
)

// GetClassificationWordPair will return a random word pair that has annotation count less than the targeted,
// but will more likely return word pair that has annotation count near the target
func GetClassificationWordPair(playerID int64) (*domain.WordPair, error) {
	annotationCriteria := domain.AnnotationCriteria{
		NotAnnotatedByPlayerID: playerID,
		MaxCount:               config.GetAppConfig().TargetAnnotationCountPerWordPair,
	}

	player, err := domain.GetPlayerByID(playerID)
	if err != nil {
		return nil, err
	}

	annotationCriteria.IsGoldStandard = shouldTrickPlayer(player.AnnotationCount)

	wordPairs, err := domain.GetWordPairByAnnotationCriteria(annotationCriteria, 5)
	if err != nil {
		return nil, err
	}

	if wordPairs == nil || len(wordPairs) == 0 {
		return nil, nil
	}

	randomIdx := rand.Intn(len(wordPairs))

	return &wordPairs[randomIdx], err
}

func shouldTrickPlayer(score int) bool {
	return score == 3 || score%7 == 0
}

func GetGoldStandards() ([]domain.GoldStandard, error) {
	goldStandards, err := domain.GetGoldStandardWordPairs()
	return goldStandards, err
}
