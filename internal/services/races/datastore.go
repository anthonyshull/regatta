package races

import "github.com/scylladb/gocqlx/v2/table"

var metadata = table.Metadata{
	Name:    "races",
	Columns: []string{"id", "name", "start", "distance"},
	PartKey: []string{"id"},
}

var tb = table.New(metadata)
