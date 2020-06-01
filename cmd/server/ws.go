package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/gorilla/websocket"
)

func upgrade(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Error().Err(err)
			return
		}

		ctx := context.WithValue(r.Context(), "conn", conn)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}


