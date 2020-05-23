package types

import "github.com/gocql/gocql"

// ID is a useful representation for a single entity passed in JSON-RPC params.
type ID struct {
	ID gocql.UUID
}
