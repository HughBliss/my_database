package example

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	xidmixin "github.com/hughbliss/my_database/internal/mixin/xid"
)

// Example holds the schema definition for the Example entity.
type Example struct {
	ent.Schema
}

// Fields of the Example.
func (Example) Fields() []ent.Field {
	return []ent.Field{
		field.String("some_string").Default("test"),
		field.Int("some_int"),
		field.Float("some_float"),
		field.Bool("some_bool"),
		field.Time("some_time").Annotations(entgql.Type("Time")),
	}
}

func (Example) Mixin() []ent.Mixin {
	return []ent.Mixin{
		xidmixin.Mixin{},
	}
}
