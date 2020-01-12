package server

import (
	"com/game_logic"
	"com/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var loggerPath = "./log.log"

func games(w http.ResponseWriter, req *http.Request) {

	switch req.Method {

	case "GET":
		keys, ok := req.URL.Query()["id"]
		if !ok || len(keys) < 1 {
			util.LogToFile(loggerPath, "missing param: 400")
			http.Error(w, "missing id param", http.StatusBadRequest)
		} else {
			key, e := strconv.Atoi(keys[0])
			if e != nil {
				util.LogToFile(loggerPath, "cannot parse param: 400")
				http.Error(w, "cannot parse param", http.StatusBadRequest)

			}
			resGame := game_logic.Get(key)
			_, e = fmt.Fprint(w, "the game was: played at: "+resGame.DateTime+" and: "+fmt.Sprint(resGame.Played))

		}

	case "POST":
		var incomingGame game_logic.IncomingGame
		err := json.NewDecoder(req.Body).Decode(&incomingGame)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		played := incomingGame.Played

		if checkPlayedSlice(played) {
			game_logic.Put(incomingGame)
			_, e := fmt.Fprintf(w, "game saved")
			if e != nil {
				util.LogToFile(loggerPath, e.Error())
			}
		} else {
			util.LogToFile(loggerPath, "unable to verify game")
			http.Error(w, "unable to verify game", http.StatusBadRequest)
		}

	default:
		util.LogToFile(loggerPath, "method invalid on endpoint")
		http.Error(w, "method invalid on endpoint", http.StatusBadRequest)
	}
}

func RunServer(port string) {
	http.HandleFunc("/games", games)
	e := http.ListenAndServe(":"+port, nil)
	if e != nil {
		util.LogToFile(loggerPath, e.Error())
	}
}
