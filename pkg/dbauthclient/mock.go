package dbauthclient

import (
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/hughbliss/my_database/pkg/gen/dbauth"
	"github.com/hughbliss/my_database/pkg/gen/dbauth/enttest"
	_ "github.com/hughbliss/my_database/pkg/gen/dbauth/runtime"
	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

func Mock(t enttest.TestingT) *dbauth.Client {
	db, err := otelsql.Open(dialect.MySQL, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	return enttest.NewClient(t, enttest.WithOptions(dbauth.Driver(sql.OpenDB(dialect.MySQL, db))))
}
