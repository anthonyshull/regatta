package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/scylladb/gocqlx/v2"
	flag "github.com/spf13/pflag"

	"github.com/anthonyshull/regatta/internal/services/races"
	"github.com/anthonyshull/regatta/internal/services/results"
	"github.com/anthonyshull/regatta/internal/services/shells"
)

func health(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("ok"))
}

func main() {
	hosts := flag.StringSlice("hosts", []string{"127.0.0.1"}, "Cassandra Hosts")
	port := flag.Int("port", 9999, "HTTP Port")
	flag.Parse()

	// logger
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

	// cassandra
	cluster := gocql.NewCluster(*hosts...)
	cluster.Keyspace = "regatta"
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		log.Fatal().Err(err)
	}

	router := mux.NewRouter()

	// middlewares

	// services
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterService(&shells.Service{Session: session}, "ShellsService")
	s.RegisterService(&races.Service{Session: session}, "RacesService")
	router.Handle("/rpc", s).Methods("POST")

	// websockets
	router.HandleFunc("/results/regatta/{id:[0-9]+}", results.Regatta).Methods("GET")

	// health
	router.HandleFunc("/health", health).Methods("GET")

	err = http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
	if err != nil {
		log.Fatal().Err(err)
	}
}
