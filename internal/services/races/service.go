package races

import (
	"net/http"

	"github.com/gocql/gocql"
	"github.com/rs/zerolog/log"
	"github.com/scylladb/gocqlx/v2"
)

type Service struct {
	gocqlx.Session
}

type Get struct {
	ID gocql.UUID
}

func (s *Service) Get(_ *http.Request, get *Get, race *Race) error {

	q := s.Session.Query(tb.Get()).BindStruct(get)
	if err := q.GetRelease(race); err != nil {
		log.Error().Err(err)
		return err
	}

	return nil
}
