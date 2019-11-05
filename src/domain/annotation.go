package domain

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
)

type Annotation struct {
	WordPairID         int   `json:"wp_id"`
	WordRelationTypeID int   `json:"wrt_id"`
	PlayerID           int64 `json:"player_id"`
	PlayerTimeMs       int   `json:"player_time_ms"`
}

func AddAnnotationAndAddPlayerScore(annotation *Annotation, score int) error {
	db := config.GetDB()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	defer cancel()

	if err != nil {
		log.Printf("DB tx initialization error: %+v\n", err)
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO annotation (
			wp_id,
			wrt_id,
			player_id,
			player_time_ms,
			created_at
		) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP)
	`, annotation.WordPairID, annotation.WordRelationTypeID, annotation.PlayerID, annotation.PlayerTimeMs)

	if err != nil {
		log.Printf("DB tx insertion error: %+v\n", err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		UPDATE player
		SET
			score = score + $1,
			annotation_count = annotation_count + 1
		WHERE
			id = $2
	`, score, annotation.PlayerID)

	if err != nil {
		log.Printf("DB tx update error: %+v\n", err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("DB tx commit error: %+v\n", err)
		tx.Rollback()
		return err
	}

	return nil
}
