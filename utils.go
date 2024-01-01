package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Utils struct{}

var utils Utils

func (Utils) Json(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (Utils) Log(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Default().Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		fn(w, r)
	}
}
