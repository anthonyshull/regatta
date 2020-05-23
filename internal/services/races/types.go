package races

import (
	"time"

	"github.com/anthonyshull/regatta/pkg/types"
)

//
type Race struct {
	types.Flyweight
	Start    time.Time
	Distance int
	Shells []types.Flyweight
}
