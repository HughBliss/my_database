package auth

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

type UserDomain struct {
	ent.Schema
}

func (UserDomain) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").GoType(xid.ID{}),
		field.String("domain_id").GoType(xid.ID{}),
		field.String("role_id").GoType(xid.ID{}),
	}
}

func (UserDomain) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("domain_id", "user_id"),
	}
}

func (UserDomain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required().Unique().Field("user_id"),
		edge.To("domain", Domain.Type).Required().Unique().Field("domain_id"),
		edge.To("role", Role.Type).Required().Unique().Field("role_id"),
	}
}

func (UserDomain) Mixin() []ent.Mixin {
	return []ent.Mixin{}
}
