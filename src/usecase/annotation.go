package usecase

import (
	"errors"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
	"github.com/yerlandinata/word-relation-gamification-backend/src/domain"
)

const (
	HyponymyRelationID  int = 1
	SynonymyRelationID  int = 2
	UnrelatedRelationID int = 3
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
				score = -2 * player.Level
				// add more penalty if the player is playing too fast
				if annotation.PlayerTimeMs < 1000 {
					score = -4 * player.Level
				}
				if annotation.PlayerTimeMs < 500 {
					score = -7 * player.Level
				}
			}
		}
	} else {
		minTime := 0
		if player.Level == 1 {
			minTime = 1000
		} else if player.Level < 5 {
			minTime = 2000
		} else {
			minTime = 4000
		}
		if annotation.PlayerTimeMs < minTime {
			score = 0
			annotation.WordRelationTypeID = config.GetAppConfig().NotSureAnnotationDBID
		}
	}

	score = score * player.Level

	err = domain.AddAnnotationAndAddPlayerScore(annotation, score)
	if err != nil {
		return nil, err
	}

	player.Score += score
	player.ElapsedTime += annotation.PlayerTimeMs

	return player, err
}

func InvalidateAnnotationsByPlayerAndGoldStandardAgreements(overallRate, perRelationTypeRate float32) ([]int64, error) {
	annotations, err := domain.GetAllAnnotations()
	if err != nil {
		return nil, err
	}

	agreements := make(map[int64]*PlayerGoldStandardAgreements)

	for _, a := range annotations {
		if _, ok := agreements[a.PlayerID]; !ok {
			agreements[a.PlayerID] = &PlayerGoldStandardAgreements{}
		}

		if a.GoldStandardRelationTypeID == 0 {
			continue
		}

		if a.WordRelationTypeID != a.GoldStandardRelationTypeID {
			agreements[a.PlayerID].OverallDisagree++
			switch g := a.GoldStandardRelationTypeID; g {
			case HyponymyRelationID:
				agreements[a.PlayerID].HyponymyDisagree++
			default:
				agreements[a.PlayerID].UnrelatedDisagree++
			}
		} else {
			agreements[a.PlayerID].OverallAgree++
			switch g := a.GoldStandardRelationTypeID; g {
			case HyponymyRelationID:
				agreements[a.PlayerID].HyponymyAgree++
			default:
				agreements[a.PlayerID].UnrelatedAgree++
			}
		}
	}

	invalidatedPlayers := make([]int64, 0)

	for k, v := range agreements {
		totalOverall := v.OverallAgree + v.OverallDisagree
		totalHyponymy := v.HyponymyAgree + v.HyponymyDisagree
		totalUnrelated := v.UnrelatedAgree + v.UnrelatedDisagree

		if totalOverall == 0 || totalHyponymy == 0 || totalUnrelated == 0 {
			continue
		}

		if float32(v.OverallAgree)/float32(totalOverall) < overallRate ||
			float32(v.HyponymyAgree)/float32(totalHyponymy) < perRelationTypeRate ||
			float32(v.UnrelatedAgree)/float32(totalUnrelated) < perRelationTypeRate {
			invalidatedPlayers = append(invalidatedPlayers, k)
		}
	}

	err = domain.InvalidateAnnotationsByPlayerIDs(invalidatedPlayers)
	if err != nil {
		return nil, err
	}

	return invalidatedPlayers, nil
}
