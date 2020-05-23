package shells

import "github.com/anthonyshull/regatta/pkg/types"

//
type Shell struct {
	types.Flyweight
	Team  types.Flyweight
	Users []types.Flyweight
}
