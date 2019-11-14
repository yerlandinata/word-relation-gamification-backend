package domain

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
)

type Annotation struct {
	WordPairID                 int   `json:"wp_id"`
	WordRelationTypeID         int   `json:"wrt_id"`
	PlayerID                   int64 `json:"player_id"`
	PlayerTimeMs               int   `json:"player_time_ms"`
	GoldStandardRelationTypeID int
}

func GetAllAnnotations() ([]Annotation, error) {
	var annotations []Annotation
	db := config.GetDB()

	rows, err := db.Query(`
		SELECT
			a.wp_id,
			a.wrt_id,
			g.wrt_id,
			a.player_id
		FROM annotation a
		LEFT JOIN gold_standard g ON g.wp_id=a.wp_id
	`)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Printf("DB query error: %+v\n", err)
		return nil, err
	}

	for rows.Next() {
		var a Annotation
		var g sql.NullInt32

		err = rows.Scan(&a.WordPairID, &a.WordRelationTypeID, &g, &a.PlayerID)
		if err != nil {
			log.Printf("DB query error: %+v\n", err)
			return nil, err
		}
		if g.Valid {
			a.GoldStandardRelationTypeID = int(g.Int32)
		}

		annotations = append(annotations, a)
	}

	return annotations, err

}

func InvalidateAnnotationsByPlayerIDs(playerIDs []int64) error {
	db := config.GetDB()

	_, err := db.Exec(`
		UPDATE annotation
		SET
			is_valid = false
		WHERE player_id = ANY($1)
	`, pq.Array(playerIDs))

	if err != nil {
		log.Printf("DB update error: %+v\n", err)
	}

	return err
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
			created_at,
			is_valid
		) VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, true)
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
			annotation_count = annotation_count + 1,
			elapsed = elapsed + $2
		WHERE
			id = $3
	`, score, annotation.PlayerTimeMs, annotation.PlayerID)

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
