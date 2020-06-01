package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

func logHTTP(next http.Handler, w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	next.ServeHTTP(w, r)

	duration := time.Now().Sub(start)

	log.
		Info().
		Str("method", r.Method).
		Str("route", r.RequestURI).
		Int64("duration", duration.Microseconds()).
		Send()
}

func logRPC(next http.Handler, w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Send()
	}

	rpc := &struct {
		Method string
	}{}

	err = json.Unmarshal(buf, rpc)
	if err != nil {
		log.Error().Err(err).Send()
	}

	r.Body = ioutil.NopCloser(bytes.NewReader(buf))

	next.ServeHTTP(w, r)

	duration := time.Now().Sub(start)

	log.
		Info().
		Str("method", r.Method).
		Str("route", r.RequestURI).
		Str("rpc", rpc.Method).
		Int64("duration", duration.Microseconds()).
		Send()
}

func telemetry(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost && r.RequestURI == "/p/rpc" {
			logRPC(next, w, r)
		}

		if r.Method == http.MethodGet {
			logHTTP(next, w, r)
		}

	})
}
