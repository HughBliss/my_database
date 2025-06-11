package auth

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	internalMixin "github.com/hughbliss/my_database/internal/mixin"
	"github.com/rs/xid"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {

	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Strings("permissions"),
		field.String("domain_id").GoType(xid.ID{}),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("domain", Domain.Type).Ref("roles").Required().Unique().Field("domain_id"),
		edge.From("user_domain", UserDomain.Type).Ref("role"),
	}
}
func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		internalMixin.XIDMixin{},
		mixin.Time{},
	}
}
