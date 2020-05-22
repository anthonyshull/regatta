package boats

import "github.com/anthonyshull/regatta/pkg/types"

//
type Boat struct {
	types.ID
	Name  string
	Team  types.Flyweight
	Users []types.Flyweight
}
