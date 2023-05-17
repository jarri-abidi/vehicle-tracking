package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-kit/log"
)

func main() {
	logger := log.NewJSONLogger(os.Stderr)
	defer logger.Log("msg", "terminated")

	addr := flag.String("addr", ":56000", "listen addr")
	flag.Parse()

	mux := http.NewServeMux()
	mux.Handle("/", NewMockServer())
	server := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	go func() {
		logger.Log("transport", "http", "address", *addr, "msg", "listening")
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Log("transport", "http", "address", *addr, "msg", "failed", "err", err)
			sig <- os.Interrupt // trigger shutdown of other resources
		}
	}()

	logger.Log("received", <-sig, "msg", "terminating")
	if err := server.Shutdown(context.Background()); err != nil {
		logger.Log("msg", "could not shutdown http mock server", "err", err)
	}
}
