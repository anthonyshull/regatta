package races

import (
	"github.com/gocql/gocql"
	"time"
)

type Race struct {
	gocql.UUID
	Name string
	Start time.Time
	Distance int
}
