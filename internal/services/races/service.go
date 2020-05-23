package races

import (
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"

	"github.com/anthonyshull/regatta/pkg/types"
)

//
type Service struct {
	gocqlx.Session
}

//
func (s *Service) Create(_ *http.Request, race *Race, _ *json2.EmptyResponse) error {
	return Create(s.Session, race)
}

//
func (s *Service) Read(_ *http.Request, id *types.ID, race *Race) error {
	return Read(s.Session, id, race)
}

//
func (s *Service) Update(_ *http.Request, race *Race, _ *json2.EmptyResponse) error {
	return Update(s.Session, race)
}

//
func (s *Service) Delete(_ *http.Request, id *types.ID, _ *json2.EmptyResponse) error {
	return Delete(s.Session, id)
}

func (s *Service) AddShell(_ *http.Request, relation *types.Relation, _ *json2.EmptyResponse) error {
	stmt, names := qb.
		Update("races").
		AddNamed("shells", "child").
		Where(qb.Eq("id")).
		ToCql()

	return s.Session.Query(stmt, names).BindStruct(relation).ExecRelease()
}

func (s *Service) RemoveShell(_ *http.Request, relation *types.Relation, _ *json2.EmptyResponse) error {
	stmt, names := qb.
		Update("races").
		RemoveNamed("shells", "child").
		Where(qb.Eq("id")).
		ToCql()
	return s.Session.Query(stmt, names).BindStruct(relation).ExecRelease()
}
