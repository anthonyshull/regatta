package results

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func Regatta(w http.ResponseWriter, r *http.Request) {
	conn := r.Context().Value("conn").(*websocket.Conn)

	vars := mux.Vars(r)

	conn.WriteJSON(vars["id"])

	conn.Close()

}