package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"

	"github.com/jarri-abidi/vehicle-tracking/karma"
)

func main() {

	addr := flag.String("addr", ":56000", "listen addr")
	flag.Parse()

	log.Printf("transport=http, address=%s, msg=listening\n", *addr)

	http.HandleFunc("/", handleRequest)
	if err := http.ListenAndServe(*addr, nil); err != http.ErrServerClosed {
		log.Printf("transport=http, address=%s, msg=failed, err=%v\n", *addr, err)
	}
}

const (
	contentTypeKey   = "Content-Type"
	contentTypeValue = "application/json; charset=utf-8"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Cmd string `json:"cmd"`
	}
	json.NewDecoder(r.Body).Decode(&request)

	var response any
	switch request.Cmd {
	case karma.CmdTrips:
		response = karma.FetchTripsResponse{Data: karma.SampleTrips(10000)}
	case karma.CmdLocation:
		response = karma.FetchLocationResponse{Data: karma.SampleLocations(10000)}
	}

	w.Header().Set(contentTypeKey, contentTypeValue)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
