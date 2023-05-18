package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":56000", "listen addr")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/trips", handleListTrips)
	mux.HandleFunc("/locations", handleListLocations)

	log.Printf("transport=http, address=%s, msg=listening\n", *addr)

	if err := http.ListenAndServe(*addr, mux); err != http.ErrServerClosed {
		log.Printf("transport=http, address=%s, msg=failed, err=%v\n", *addr, err)
	}
}
