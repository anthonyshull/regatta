package boats

import (
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
	"github.com/rs/zerolog/log"
	"github.com/scylladb/gocqlx/v2"

	"github.com/anthonyshull/regatta/pkg/types"
)

//
type Service struct {
	gocqlx.Session
}

//
func (s *Service) Create(_ *http.Request, boat *Boat, _ *json2.EmptyResponse) error {

	q := s.Session.Query(tb.Insert()).BindStruct(boat)
	if err := q.Exec(); err != nil {
		log.Error().Err(err)
		return err
	}

	return nil
}

//
func (s *Service) Read(_ *http.Request, id *types.ID, boat *Boat) error {

	q := s.Session.Query(tb.Get()).BindStruct(id)
	if err := q.GetRelease(boat); err != nil {
		log.Error().Err(err)
		return err
	}

	return nil
}

//
func (s *Service) Delete(_ *http.Request, id *types.ID, _ *json2.EmptyResponse) error {

	q := s.Session.Query(tb.Delete()).BindStruct(id)
	if err := q.Exec(); err != nil {
		log.Error().Err(err)
		return err
	}

	return nil
}
