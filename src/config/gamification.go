package config

import (
	"log"
	"os"
	"strconv"
)

type AppConfig struct {
	TargetAnnotationCountPerWordPair int
	NotSureAnnotationDBID            int
	Secret                           string
}

var config *AppConfig

func InitAppConfig() {
	target, err := strconv.Atoi(os.Getenv("TARGET_ANNOTATION_COUNT_PER_WORD_PAIR"))
	if err != nil {
		log.Fatal("Failed to load gamification config!")
	}

	notSureID, err := strconv.Atoi(os.Getenv("NOTSURE_WRT_ID"))
	if err != nil {
		log.Fatal("Failed to load gamification config!")
	}

	config = &AppConfig{
		TargetAnnotationCountPerWordPair: target,
		Secret:                           os.Getenv("SECRET"),
		NotSureAnnotationDBID:            notSureID,
	}
}

func GetAppConfig() *AppConfig {
	if config == nil {
		InitAppConfig()
	}
	return config
}
