package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yerlandinata/word-relation-gamification-backend/src/server/httphandler"
	"github.com/yerlandinata/word-relation-gamification-backend/src/utils/httputils"
)

func Serve() {
	http.HandleFunc("/", helloServer)
	httputils.HandlePost("/users/login", httphandler.Login)
	httputils.HandlePost("/users/register", httphandler.Register)
	httputils.HandlePost("/users/restart_game", httputils.Authenticate(httphandler.ResetPlayerScoreAndTimeAndContinueGame))
	httputils.HandleGet("/users/word_pairs", httputils.Authenticate(httphandler.GetClassificationWordPair))
	httputils.HandlePost("/users/word_pairs/annotations", httputils.Authenticate(httphandler.AddAnnotation))
	httputils.HandleGet("/gold_standards", httphandler.GetGoldStandards)

	port := os.Getenv("PORT")
	log.Printf("listening 0.0.0.0:%s", port)
	err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), nil)
	log.Printf("HTTP Web server fails: %+v\n", err)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	httputils.ResponseJSON(w, http.StatusNotFound, struct{ Message string }{Message: "Not found"})
}
