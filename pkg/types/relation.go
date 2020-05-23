package types

import "github.com/gocql/gocql"

// Relation aids in adding or removing elements from an entity with child relations (list, set).
type Relation struct {
	ID gocql.UUID
	Child []Flyweight
}
