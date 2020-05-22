package races

import (
	"time"

	"github.com/anthonyshull/regatta/pkg/types"
)

//
type Race struct {
	types.ID
	Name     string
	Start    time.Time
	Distance int
}
