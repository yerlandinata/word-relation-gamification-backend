package domain

import (
	"database/sql"
	"log"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
)

type Player struct {
	ID              int64  `json:"id"`
	FullName        string `json:"full_name"`
	Password        int64  `json:"birth_date"`
	EducationLevel  string `json:"education_level"`
	Score           int64  `json:"score"`
	AnnotationCount int    `json:"annotation_count"`
	Rank            int64  `json:"rank"`
	ElapsedTime     int    `json:"elapsed"`
}

func GetPlayerByID(id int64) (*Player, error) {
	var result Player
	db := config.GetDB()

	row := db.QueryRow(`
		SELECT 
			p1.id,
			p1.full_name,
			p1.birth_date,
			p1.score,
			p1.annotation_count,
			p1.ranking,
			p1.elapsed
		FROM (
			SELECT 
				p2.id,
				p2.full_name,
				p2.birth_date,
				p2.score,
				p2.annotation_count,
				p2.elapsed,
				RANK () OVER (
					ORDER BY p2.score DESC
				) ranking
			FROM player p2
		) p1
		WHERE p1.id=$1
	`, id)

	err := row.Scan(&result.ID, &result.FullName, &result.Password, &result.Score, &result.AnnotationCount, &result.Rank, &result.ElapsedTime)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Printf("DB query error: %+v\n", err)
	}

	return &result, err
}

func AddPlayer(player *Player) error {
	db := config.GetDB()

	_, err := db.Exec(`
		INSERT INTO player (
			id,
			full_name,
			birth_date,
			education_level,
			score,
			annotation_count,
			elapsed
		) VALUES ($1, $2, $3, $4, 0, 0, 0)
	`, player.ID, player.FullName, player.Password, player.EducationLevel)

	if err != nil {
		log.Printf("DB insertion error: %+v\n", err)
	}

	return err
}

func ResetPlayerScoreAndTime(playerID int64) error {
	db := config.GetDB()

	_, err := db.Exec(`
		UPDATE player
		SET
			score = 0,
			elapsed = 0
		WHERE id = $1
	`, playerID)

	if err != nil {
		log.Printf("DB update error: %+v\n", err)
	}

	return err
}
