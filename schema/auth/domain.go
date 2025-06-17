package auth

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	xidmixin "github.com/hughbliss/my_database/internal/mixin/xid"
)

// Domain holds the schema definition for the Domain entity.
type Domain struct {
	ent.Schema
}

// Fields of the Domain.
func (Domain) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		//field.JSON("settings", map[string]interface{}{}),
	}
}

// Edges of the Domain.
func (Domain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("active_users", User.Type),
		edge.To("roles", Role.Type),

		edge.To("users", User.Type).Through("user_domain", UserDomain.Type),
	}
}

func (Domain) Mixin() []ent.Mixin {
	return []ent.Mixin{
		xidmixin.Mixin{},
	}
}
