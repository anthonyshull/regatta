package shells

import (
	"net/http"

	"github.com/gorilla/rpc/v2/json2"
	"github.com/scylladb/gocqlx/v2"

	"github.com/anthonyshull/regatta/pkg/types"
)

//
type Service struct {
	gocqlx.Session
}

//
func (s *Service) Create(_ *http.Request, shell *Shell, _ *json2.EmptyResponse) error {
	return Create(s.Session, shell)
}

//
func (s *Service) Read(_ *http.Request, id *types.ID, shell *Shell) error {
	return Read(s.Session, id, shell)
}

//
func (s *Service) Update(_ *http.Request, shell *Shell) error {
	return Update(s.Session, shell)
}

//
func (s *Service) Delete(_ *http.Request, id *types.ID, _ *json2.EmptyResponse) error {
	return Delete(s.Session, id)
}
