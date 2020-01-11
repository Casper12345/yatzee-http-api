package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Game struct {
	DateTime string
	Played   [][] int
}

func hello(w http.ResponseWriter, req *http.Request) {

	f := req.Header

	//return "hi there"
	method := req.Method

	switch method {
	case "GET":
		fmt.Fprintf(w, "im get\n"+f.Get("cookie"))
	case "POST":
		var g Game
		err := json.NewDecoder(req.Body).Decode(&g)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		played := g.Played

		if checkPlayedSlice(played) {
			fmt.Fprintf(w, "cool")
		} else {
			http.Error(w, "damn", http.StatusBadRequest)
		}

	default:
		fmt.Fprintf(w, "sorry")
	}
}

func checkPlayedSlice(s [][] int) bool {
	if len(s) != 13 {
		return false
	} else {
		for _, i := range s {
			if len(i) != 5 {
				return false
			} else {
				for _, i2 := range i {
					if i2 > 6 || i2 < 1 {
						return false
					}
				}
			}
		}
		return true
	}
}

func another(w http.ResponseWriter, req *http.Request){
	http.Error(w, "jeez", http.StatusBadRequest)
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func Mother() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/header", headers)
	http.HandleFunc("/another", another)

	http.ListenAndServe(":8090", nil)

}
