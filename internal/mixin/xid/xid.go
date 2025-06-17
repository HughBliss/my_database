package xidmixin

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/rs/xid"
)

type Mixin struct {
	mixin.Schema
}

func (Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(xid.ID{}).
			SchemaType(map[string]string{
				dialect.Postgres: "char(20)",
			}).
			DefaultFunc(xid.New),
	}
}
