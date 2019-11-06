package config

import (
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	TargetAnnotationCountPerWordPair     int
	TargetAnnotationCountPerGoldStandard int
	NotSureAnnotationDBID                int
	Secret                               string
}

var config *AppConfig

func InitAppConfig() {
	target, err := strconv.Atoi(os.Getenv("TARGET_ANNOTATION_COUNT_PER_WORD_PAIR"))
	if err != nil {
		log.Fatal("Failed to load gamification config!")
	}

	targetGold, err := strconv.Atoi(os.Getenv("TARGET_ANNOTATION_COUNT_PER_GOLD_STANDARD"))
	if err != nil {
		log.Fatal("Failed to load gamification config!")
	}

	notSureID, err := strconv.Atoi(os.Getenv("NOTSURE_WRT_ID"))
	if err != nil {
		log.Fatal("Failed to load gamification config!")
	}

	config = &AppConfig{
		TargetAnnotationCountPerWordPair:     target,
		TargetAnnotationCountPerGoldStandard: targetGold,
		Secret:                               os.Getenv("SECRET"),
		NotSureAnnotationDBID:                notSureID,
	}
}

func GetAppConfig() *AppConfig {
	if config == nil {
		InitAppConfig()
	}
	return config
}
