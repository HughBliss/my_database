package example

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/hughbliss/my_database/internal/mixin"
)

// Example holds the schema definition for the Example entity.
type Example struct {
	ent.Schema
}

// Fields of the Example.
func (Example) Fields() []ent.Field {
	return []ent.Field{
		field.String("some_string"),
		field.Int("some_int"),
		field.Float("some_float"),
		field.Bool("some_bool"),
		field.Time("some_time"),
	}
}

func (Example) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.XIDMixin{},
	}
}
