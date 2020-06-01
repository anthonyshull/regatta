package shells

import (
	"github.com/gocql/gocql"

	"github.com/anthonyshull/regatta/pkg/types"
)

//
type Shell struct {
	ID    gocql.UUID
	Name  string
	Team  types.Flyweight
	Users []types.Flyweight
}
