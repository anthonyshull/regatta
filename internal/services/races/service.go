package races

import (
	"net/http"

	"github.com/gocql/gocql"
	"github.com/rs/zerolog"
	"github.com/scylladb/gocqlx/v2"
)

type Service struct {
	gocqlx.Session
	zerolog.Logger
}

type Get struct {
	gocql.UUID
}

func (s *Service) Get(_ *http.Request, get *Get, race *Race) error {
	race.UUID = get.UUID
	q := s.Session.Query(tb.Get()).BindStruct(race)
	if err := q.GetRelease(&race); err != nil {
		s.Logger.Error().Str("I'M", "HERE")
		return err
	}
	return nil
}
