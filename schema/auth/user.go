package auth

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	internalMixin "github.com/hughbliss/my_database/internal/mixin"
	"github.com/rs/xid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.String("password_hash"),
		field.String("current_domain_id").GoType(xid.ID{}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("current_domain", Domain.Type).Ref("active_users").Unique().Required().Field("current_domain_id"),
		edge.From("domains", Domain.Type).Ref("users").Through("user_domain", UserDomain.Type),
	}
}
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		internalMixin.XIDMixin{},

		mixin.Time{},
	}
}
