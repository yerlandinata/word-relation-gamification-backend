package server

import (
	"log"

	"github.com/yerlandinata/word-relation-gamification-backend/src/usecase"
)

var tasks map[string]func()

func RegisterTasks() {
	tasks = make(map[string]func())
	tasks["filter-annotators"] = filterAnnotators
}

func RunTask(taskID string) {
	if taskFunc, ok := tasks[taskID]; ok {
		taskFunc()
	}
}

func filterAnnotators() {

	log.Println("filtering annotators..")

	playerIDs, err := usecase.InvalidateAnnotationsByPlayerAndGoldStandardAgreements(.5, .4)
	if err != nil {
		log.Println(err)
	}

	log.Println("invalidated annotators:", playerIDs)
}
