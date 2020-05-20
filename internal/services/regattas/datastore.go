package regattas

import "github.com/scylladb/gocqlx/v2/table"

var metadata = table.Metadata{
	Name:    "regattas",
	Columns: []string{"id", "name", "start", "stop", "races"},
	PartKey: []string{"id"},
	SortKey: []string{"id"},
}