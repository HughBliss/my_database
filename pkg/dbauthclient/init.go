package dbauthclient

import (
	"context"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"fmt"
	zfg "github.com/chaindead/zerocfg"
	"github.com/hughbliss/my_database/pkg/gen/dbauth"
	_ "github.com/hughbliss/my_database/pkg/gen/dbauth/runtime"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"

	"github.com/uptrace/opentelemetry-go-extra/otelsql"
)

var (
	group    = zfg.NewGroup("db_auth")
	host     = zfg.Str("host", "", "DBAUTH_HOST", zfg.Required(), zfg.Group(group))
	name     = zfg.Str("name", "", "DBAUTH_NAME", zfg.Required(), zfg.Group(group))
	user     = zfg.Str("user", "", "DBAUTH_USER", zfg.Required(), zfg.Group(group))
	port     = zfg.Uint32("port", 0, "DBAUTH_PORT", zfg.Required(), zfg.Group(group))
	password = zfg.Str("password", "", "DBAUTH_PASSWORD", zfg.Required(), zfg.Group(group), zfg.Secret())
	sslmode  = zfg.Str("ssl_mode", "disable", "DBAUTH_SSLMODE", zfg.Group(group))
)

type Config struct {
	DisableAutoMigration bool
	Debug                bool
}

func Init(config *Config) (client *dbauth.Client, err error) {
	if config == nil {
		config = &Config{}
	}
	db, err := otelsql.Open(dialect.Postgres, fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s", *host, *port, *user, *name, *password, *sslmode))
	if err != nil {
		return nil, err
	}

	client = dbauth.NewClient(dbauth.Driver(sql.OpenDB(dialect.Postgres, db)))

	if config.Debug {
		client = client.Debug()
	}

	if !config.DisableAutoMigration {
		if err := client.Schema.Create(context.Background()); err != nil {
			return nil, errors.Wrap(err, "error creating database schema")
		}
	}

	return client, nil

}
