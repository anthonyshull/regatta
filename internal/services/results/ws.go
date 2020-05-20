package results

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func Regatta(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	vars := mux.Vars(r)
	conn.WriteJSON(vars["id"])
	conn.Close()
}