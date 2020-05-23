package shells

import (
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"

	"github.com/anthonyshull/regatta/pkg/types"
)

var metadata = table.Metadata{
	Name:    "shells",
	Columns: []string{"id", "name", "team", "users"},
	PartKey: []string{"id"},
}

var tb = table.New(metadata)

func Create(session gocqlx.Session, shell *Shell) error {
	q := session.Query(tb.Insert()).BindStruct(shell)
	return q.ExecRelease()
}

func Read(session gocqlx.Session, id *types.ID, shell *Shell) error {
	q := session.Query(tb.Get()).BindStruct(id)
	return q.GetRelease(shell)
}

func Update(session gocqlx.Session, shell *Shell) error {
	q := session.Query(tb.Update()).BindStruct(shell)
	return q.ExecRelease()
}

func Delete(session gocqlx.Session, id *types.ID) error {
	q := session.Query(tb.Delete()).BindStruct(id)
	return q.ExecRelease()
}