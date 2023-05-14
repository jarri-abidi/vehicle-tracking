package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":56000", "listen addr")
	flag.Parse()

	if err := http.ListenAndServe(*addr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("karma/trips_sample.json")
		if err != nil {
			log.Printf("could not read file: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		_, err = io.Copy(w, f)
		if err != nil {
			log.Printf("could not write to response writer: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})); err != nil {
		log.Fatal("could not serve http: ", err)
	}
}
