package usecase

import (
	"errors"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
	"github.com/yerlandinata/word-relation-gamification-backend/src/domain"
)

// AddAnnotation will add annotation to our DB,
// If the word pair is already in gold standard: if correct get 2 score, if wrong get -1 score
// If the word pair is not in gold standard player will get +1 score
func AddAnnotation(annotation *domain.Annotation) (*domain.Player, error) {

	player, err := domain.GetPlayerByID(annotation.PlayerID)
	if err != nil {
		return nil, err
	}

	if player.ElapsedTime > config.GetAppConfig().GameTimeLimitMS {
		return nil, errors.New("Time's up")
	}

	score := 0

	if annotation.WordRelationTypeID != config.GetAppConfig().NotSureAnnotationDBID {
		score = 1
	}

	// get gold standard
	goldStandard, err := domain.GetGoldStandardByWordPairID(annotation.WordPairID)
	if err != nil {
		return nil, err
	}

	if goldStandard != nil {
		if annotation.WordRelationTypeID == goldStandard.WordRelationType.ID {
			score = 2
		} else {
			if annotation.WordRelationTypeID == config.GetAppConfig().NotSureAnnotationDBID {
				score = 0
			} else {
				score = -2
			}
		}
	}

	err = domain.AddAnnotationAndAddPlayerScore(annotation, score)
	if err != nil {
		return nil, err
	}

	player, err = domain.GetPlayerByID(annotation.PlayerID)
	if err != nil {
		return nil, err
	}

	return player, err
}
