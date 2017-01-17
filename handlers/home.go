package handlers

import (
	"../gateways/champion"
	"log"
	"net/http"
)

func Home(rw http.ResponseWriter, rq *http.Request) {
	log.Println("Executing the home handler")
	briefChampionData := make(map[string]interface{})
	briefChampionData["champions"] = champion.GetAllChampions()
	sendResponse(rw, rq, "home", briefChampionData)
}
