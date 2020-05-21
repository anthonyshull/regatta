package boats

import "github.com/scylladb/gocqlx/v2/table"

var metadata = table.Metadata{
	Name:    "boats",
	Columns: []string{"id", "name", "team", "users"},
	PartKey: []string{"id"},
}

var tb = table.New(metadata)