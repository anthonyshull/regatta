package regattas

import (
	"time"

	"github.com/anthonyshull/regatta/internal/services/races"
)

type Regatta struct {
	ID string
	Name string
	Start time.Time
	Stop time.Time
	Races []races.Race
}