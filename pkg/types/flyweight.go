package types

import (
	"reflect"

	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

// Flyweight helps to reference an entity without the need for including all fields.
type Flyweight struct {
	ID gocql.UUID
	Name string
}

// MarshalUDT marshals the Flyweight since it is a UDT in Cassandra.
func (f *Flyweight) MarshalUDT(name string, info gocql.TypeInfo) ([]byte, error) {
	field := gocqlx.DefaultMapper.FieldByName(reflect.ValueOf(f), name)
	return gocql.Marshal(info, field.Interface())
}

// UnmarshalUDT unmarshals the Flyweight since it is a UDT in Cassandra.
func (f *Flyweight) UnmarshalUDT(name string, info gocql.TypeInfo, data []byte) error {
	field := gocqlx.DefaultMapper.FieldByName(reflect.ValueOf(f), name)
	return gocql.Unmarshal(info, data, field.Addr().Interface())
}