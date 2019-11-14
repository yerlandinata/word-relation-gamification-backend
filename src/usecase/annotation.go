package usecase

import (
	"errors"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
	"github.com/yerlandinata/word-relation-gamification-backend/src/domain"
)

// AddAnnotation will add annotation to our DB and decide the score for the player
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
				// add more penalty if the player is playing too fast
				if annotation.PlayerTimeMs < 1000 {
					score = -7
				}
				if annotation.PlayerTimeMs < 500 {
					score = -20
				}
			}
		}
	} else {
		if annotation.PlayerTimeMs < 1000 {
			// too fast
			return player, nil
		}
	}

	score = score * player.Level

	err = domain.AddAnnotationAndAddPlayerScore(annotation, score)
	if err != nil {
		return nil, err
	}
	
	player.Score += score

	return player, err
}
