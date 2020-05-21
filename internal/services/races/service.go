package races

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
func (s *Service) Create(_ *http.Request, race *Race, _ *json2.EmptyResponse) error {

	q := s.Session.Query(tb.Insert()).BindStruct(race)
	if err := q.Exec(); err != nil {
		log.Error().Err(err)
		return err
	}

	return nil
}

//
func (s *Service) Read(_ *http.Request, id *types.ID, race *Race) error {

	q := s.Session.Query(tb.Get()).BindStruct(id)
	if err := q.GetRelease(race); err != nil {
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
