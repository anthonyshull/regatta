package races

import (
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/table"

	"github.com/anthonyshull/regatta/pkg/types"
)

var metadata = table.Metadata{
	Name:    "races",
	Columns: []string{"id", "name", "start", "distance", "shells"},
	PartKey: []string{"id"},
}

var tb = table.New(metadata)

func Create(session gocqlx.Session, race *Race) error {
	q := session.Query(tb.Insert()).BindStruct(race)
	return q.ExecRelease()
}

func Read(session gocqlx.Session, id *types.ID, race *Race) error {
	q := session.Query(tb.Get()).BindStruct(id)
	return q.Get(race)
}

func Update(session gocqlx.Session, race *Race) error {
	q := session.Query(tb.Update()).BindStruct(race)
	return q.ExecRelease()
}

func Delete(session gocqlx.Session, id *types.ID) error {
	q := session.Query(tb.Delete()).BindStruct(id)
	return q.ExecRelease()
}
