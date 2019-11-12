package domain

import (
	"database/sql"
	"log"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
)

type Player struct {
	ID              int64  `json:"id"`
	DisplayName     string `json:"display_name"`
	CampaignSource  string `json:"campaign_source"`
	Score           int64  `json:"score"`
	AnnotationCount int    `json:"annotation_count"`
	Rank            int    `json:"rank"`
	ElapsedTime     int    `json:"elapsed"`
	Level           int    `json:"level"`
}

func GetPlayerByID(id int64) (*Player, error) {
	var result Player
	db := config.GetDB()

	row := db.QueryRow(`
		SELECT 
			p1.id,
			p1.display_name,
			p1.score,
			p1.annotation_count,
			p1.ranking,
			p1.elapsed,
			p1.game_level
		FROM (
			SELECT 
				p2.id,
				p2.display_name,
				p2.birth_date,
				p2.score,
				p2.annotation_count,
				p2.elapsed,
				p2.game_level,
				RANK () OVER (
					ORDER BY p2.score DESC
				) ranking
			FROM player p2
		) p1
		WHERE p1.id=$1
	`, id)

	err := row.Scan(&result.ID, &result.DisplayName, &result.Score, &result.AnnotationCount, &result.Rank, &result.ElapsedTime, &result.Level)

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
			display_name,
			campaign_source,
			score,
			annotation_count,
			elapsed,
			game_level
		) VALUES ($1, $2, $3, 0, 0, 0, 1)
	`, player.ID, player.DisplayName, player.CampaignSource)

	if err != nil {
		log.Printf("DB insertion error: %+v\n", err)
	}

	return err
}

func UpdatePlayerID(player *Player, newID int64) error {
	db := config.GetDB()

	_, err := db.Exec(`
		UPDATE player
		SET
			id=$1
		WHERE id=$2
	`, newID, player.ID)

	if err != nil {
		log.Printf("DB update error: %+v\n", err)
	}

	return err

}

func IncrementPlayerLevelAndResetPlayerTime(playerID int64) error {
	db := config.GetDB()

	_, err := db.Exec(`
		UPDATE player
		SET
			elapsed = 0,
			game_level = game_level + 1
		WHERE id = $1
	`, playerID)

	if err != nil {
		log.Printf("DB update error: %+v\n", err)
	}

	return err
}

func GetPlayerRankingsRange(minRank, maxRank int) ([]Player, error) {
	var result []Player
	db := config.GetDB()

	rows, err := db.Query(`
		SELECT *
		FROM (
			SELECT
				display_name,
				score,
				game_level,
				RANK () OVER (
					ORDER BY score DESC
				) ranking
			FROM player
		) t
		WHERE ranking >= $1 AND ranking <= $2
		ORDER BY ranking ASC
	`, minRank, maxRank)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p Player

		err = rows.Scan(&p.DisplayName, &p.Score, &p.Level, &p.Rank)
		if err != nil {
			log.Printf("DB query error: %+v\n", err)
			return nil, err
		}

		result = append(result, p)
	}

	return result, nil
}
