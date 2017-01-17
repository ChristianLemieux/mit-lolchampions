package handlers

import (
	"../gateways/champion"
	"log"
	"net/http"
	"strings"
)

func Champion(rw http.ResponseWriter, rq *http.Request) {
	log.Println("Executing the champion handler")
	// Here I'm manually parsing the url and grabbing the various
	// parts. There are libraries that do a better job of this,
	// but if you are doing something really basic, this is fine
	pathParts := strings.Split(rq.URL.Path, "/")
	pathPartCount := len(pathParts)
	if pathPartCount < 3 {
		// Convenient...
		http.NotFound(rw, rq)
	}
	championId := pathParts[2]

	currentChampion := champion.GetChampion(championId)
	if currentChampion == nil {
		http.NotFound(rw, rq)
		return
	}
	// Fetch all the data and prepare it for transformation with
	// the templates.
	championData := make(map[string]interface{})
	championData["champion"] = currentChampion

	sendResponse(rw, rq, "champion", championData)
}
