package races

import "github.com/scylladb/gocqlx/v2/table"

var metadata = table.Metadata{
	Name:    "races",
	Columns: []string{"uuid", "name", "start", "distance"},
}

var tb = table.New(metadata)