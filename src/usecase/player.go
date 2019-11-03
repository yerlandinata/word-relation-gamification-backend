package usecase

import (
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
