// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldCurrentDomainID holds the string denoting the current_domain_id field in the database.
	FieldCurrentDomainID = "current_domain_id"
	// EdgeCurrentDomain holds the string denoting the current_domain edge name in mutations.
	EdgeCurrentDomain = "current_domain"
	// EdgeDomains holds the string denoting the domains edge name in mutations.
	EdgeDomains = "domains"
	// EdgeUserDomain holds the string denoting the user_domain edge name in mutations.
	EdgeUserDomain = "user_domain"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CurrentDomainTable is the table that holds the current_domain relation/edge.
	CurrentDomainTable = "users"
	// CurrentDomainInverseTable is the table name for the Domain entity.
	// It exists in this package in order to avoid circular dependency with the "domain" package.
	CurrentDomainInverseTable = "domains"
	// CurrentDomainColumn is the table column denoting the current_domain relation/edge.
	CurrentDomainColumn = "current_domain_id"
	// DomainsTable is the table that holds the domains relation/edge. The primary key declared below.
	DomainsTable = "user_domains"
	// DomainsInverseTable is the table name for the Domain entity.
	// It exists in this package in order to avoid circular dependency with the "domain" package.
	DomainsInverseTable = "domains"
	// UserDomainTable is the table that holds the user_domain relation/edge.
	UserDomainTable = "user_domains"
	// UserDomainInverseTable is the table name for the UserDomain entity.
	// It exists in this package in order to avoid circular dependency with the "userdomain" package.
	UserDomainInverseTable = "user_domains"
	// UserDomainColumn is the table column denoting the user_domain relation/edge.
	UserDomainColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldName,
	FieldEmail,
	FieldPasswordHash,
	FieldCurrentDomainID,
}

var (
	// DomainsPrimaryKey and DomainsColumn2 are the table columns denoting the
	// primary key for the domains relation (M2M).
	DomainsPrimaryKey = []string{"domain_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() xid.ID
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreateTime orders the results by the create_time field.
func ByCreateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateTime, opts...).ToFunc()
}

// ByUpdateTime orders the results by the update_time field.
func ByUpdateTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdateTime, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPasswordHash orders the results by the password_hash field.
func ByPasswordHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

// ByCurrentDomainID orders the results by the current_domain_id field.
func ByCurrentDomainID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrentDomainID, opts...).ToFunc()
}

// ByCurrentDomainField orders the results by current_domain field.
func ByCurrentDomainField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCurrentDomainStep(), sql.OrderByField(field, opts...))
	}
}

// ByDomainsCount orders the results by domains count.
func ByDomainsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDomainsStep(), opts...)
	}
}

// ByDomains orders the results by domains terms.
func ByDomains(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDomainsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserDomainCount orders the results by user_domain count.
func ByUserDomainCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserDomainStep(), opts...)
	}
}

// ByUserDomain orders the results by user_domain terms.
func ByUserDomain(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserDomainStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCurrentDomainStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CurrentDomainInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CurrentDomainTable, CurrentDomainColumn),
	)
}
func newDomainsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DomainsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, DomainsTable, DomainsPrimaryKey...),
	)
}
func newUserDomainStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserDomainInverseTable, UserDomainColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, UserDomainTable, UserDomainColumn),
	)
}
