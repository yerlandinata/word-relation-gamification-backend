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

	if player.ElapsedTime > config.GetAppConfig().GameTimeLimitMS {
		return nil, nil
	}

	annotationCriteria.IsGoldStandard = shouldTrickPlayer(player.AnnotationCount)

	if annotationCriteria.IsGoldStandard {
		annotationCriteria.MaxCount = config.GetAppConfig().TargetAnnotationCountPerGoldStandard
	}

	wordPairs, err := domain.GetWordPairByAnnotationCriteria(annotationCriteria, 20*player.Level)
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
	return score < 3 || (score%5 == 0 && score <= 30) || (score%7 == 0 && score > 30)
}

func GetGoldStandards() ([]domain.GoldStandard, error) {
	hyponymyGoldStandards, err := domain.GetGoldStandardWordPairs(HyponymyRelationID, 20)
	if err != nil {
		return nil, err
	}
	synonymyGoldStandards, err := domain.GetGoldStandardWordPairs(SynonymyRelationID, 5)
	if err != nil {
		return nil, err
	}
	unrelatedGoldStandards, err := domain.GetGoldStandardWordPairs(UnrelatedRelationID, 20)
	if err != nil {
		return nil, err
	}

	goldStandards := make([]domain.GoldStandard, 0)

	randomIdx := rand.Intn(len(hyponymyGoldStandards))
	for i := 0; i < 3; i++ {
		goldStandards = append(goldStandards, hyponymyGoldStandards[(i+randomIdx)%len(hyponymyGoldStandards)])
	}

	goldStandards = append(goldStandards, synonymyGoldStandards[0])

	randomIdx = rand.Intn(len(unrelatedGoldStandards))
	for i := 0; i < 3; i++ {
		goldStandards = append(goldStandards, unrelatedGoldStandards[(i+randomIdx)%len(unrelatedGoldStandards)])
	}
	return goldStandards, err
}
