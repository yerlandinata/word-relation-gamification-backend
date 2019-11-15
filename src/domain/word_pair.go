package domain

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/yerlandinata/word-relation-gamification-backend/src/config"
)

type GoldStandard struct {
	WordRelationType WordRelationType `json:"wrt"`
	WordPair         WordPair         `json:"wp"`
}

type WordRelationType struct {
	ID        int    `json:"id"`
	ShortDesc string `json:"short_desc"`
}

type WordPair struct {
	ID    int64  `json:"id"`
	Word1 string `json:"word_1"`
	Word2 string `json:"word_2"`
}

const (
	ActiveWordPair   int = 1
	InactiveWordPair int = 0
)

type AnnotationCriteria struct {
	MaxCount               int
	NotAnnotatedByPlayerID int64
	Word1                  string
	Word2                  string
	IsGoldStandard         bool
}

func GetWordPairByAnnotationCriteria(criteria AnnotationCriteria, limit int) ([]WordPair, error) {
	var result []WordPair

	if criteria.NotAnnotatedByPlayerID == 0 {
		return nil, errors.New("NotAnnotatedByPlayerID cannot be empty")
	}

	db := config.GetDB()

	whereStmt := `
		wp.active_status=$1
		AND a.is_valid=true
		AND wp.id NOT IN (
			SELECT wp_id FROM annotation WHERE player_id=$2
		)
	`
	params := make([]interface{}, 0)
	params = append(params, ActiveWordPair)
	params = append(params, criteria.NotAnnotatedByPlayerID)

	joinStmt := ""

	if !criteria.IsGoldStandard {
		whereStmt += `
		AND wp.id NOT IN (
			SELECT wp_id FROM gold_standard
		)		
		`
	} else {
		joinStmt = "INNER JOIN gold_standard g on g.wp_id=wp.id"
	}

	havingStmt := ""

	if criteria.Word1 != "" {
		params = append(params, criteria.Word1)
		whereStmt += fmt.Sprintf(`%s
			AND wp.word_1=$%d
		`, whereStmt, len(params))
	}

	if criteria.Word2 != "" {
		params = append(params, criteria.Word2)
		whereStmt += fmt.Sprintf(`%s
			AND wp.word_2=$%d
		`, whereStmt, len(params))
	}

	if criteria.MaxCount != 0 {
		params = append(params, criteria.MaxCount)
		havingStmt += fmt.Sprintf(`
			HAVING COUNT(a.wp_id) < $%d
		`, len(params))
	}

	query := `
		SELECT 
			wp.id,
			wp.word_1,
			wp.word_2
		FROM word_pair wp
		LEFT JOIN annotation a on a.wp_id=wp.id
	` + joinStmt + `
		WHERE
	` + whereStmt + `
		GROUP BY a.wp_id, wp.id
	` + havingStmt + `
		ORDER BY COUNT(a.wp_id) DESC, wp.pmi, wp.word_1_freq DESC
		LIMIT ` + strconv.Itoa(limit)

	rows, err := db.Query(query, params...)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Printf("DB query error: %+v\n", err)
		return nil, err
	}

	for rows.Next() {
		var wp WordPair

		err = rows.Scan(&wp.ID, &wp.Word1, &wp.Word2)
		if err != nil {
			log.Printf("DB query error: %+v\n", err)
			return nil, err
		}

		result = append(result, wp)
	}

	return result, err
}

func GetGoldStandardWordPairs(wrtID, limit int) ([]GoldStandard, error) {
	var result []GoldStandard

	db := config.GetDB()

	rows, err := db.Query(`
		SELECT wp.id, wp.word_1, wp.word_2, wrt.id, wrt.short_desc
		FROM gold_standard g
		LEFT JOIN word_pair wp ON g.wp_id=wp.id
		LEFT JOIN word_relation_type wrt ON g.wrt_id=wrt.id
		WHERE g.wrt_id=$1
		LIMIT $2
	`, wrtID, limit)

	if err != nil {
		log.Printf("DB query error: %+v\n", err)
		return nil, err
	}

	for rows.Next() {
		var g GoldStandard

		err = rows.Scan(&g.WordPair.ID, &g.WordPair.Word1, &g.WordPair.Word2, &g.WordRelationType.ID, &g.WordRelationType.ShortDesc)
		if err != nil {
			log.Printf("DB query error: %+v\n", err)
			return nil, err
		}

		result = append(result, g)
	}

	return result, nil
}

func GetGoldStandardByWordPairID(wpID int) (*GoldStandard, error) {
	var result GoldStandard
	db := config.GetDB()

	row := db.QueryRow(`
		SELECT wrt_id
		FROM gold_standard
		WHERE wp_id=$1
	`, wpID)

	err := row.Scan(&result.WordRelationType.ID)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Printf("DB query error: %+v\n", err)
	}

	return &result, err
}
