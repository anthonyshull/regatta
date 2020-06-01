package races

import (
	"time"

	"github.com/gocql/gocql"

	"github.com/anthonyshull/regatta/pkg/types"
)

//
type Race struct {
	ID       gocql.UUID
	Name     string
	Start    time.Time
	Distance int
	Shells   []types.Flyweight
}
