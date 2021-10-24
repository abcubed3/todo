package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func writeTodoJSON(w http.ResponseWriter, data *Todo) {
	// json byte
	payload, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(payload)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
