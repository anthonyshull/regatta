package races

import (
	"time"

	"github.com/gocql/gocql"
)

type Race struct {
	ID       gocql.UUID
	Name     string
	Start    time.Time
	Distance int
}
