package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	zfg "github.com/chaindead/zerocfg"
	"github.com/hughbliss/my_toolkit/cfg"
)

var (
	schemaName = zfg.Str("schema_name", "", "SCHEMANAME", zfg.Alias("s"), zfg.Required())
)

func main() {
	if err := cfg.Init(); err != nil {
		panic(err)
	}

	if err := entc.Generate( /*Schema path*/ "./schema/"+*schemaName, &gen.Config{
		/*Schema package name*/ Schema: "github.com/hughbliss/my_database/schema/" + *schemaName,
		/*Gen path*/ Target: "./pkg/gen/db" + *schemaName,
		/*Gen package name*/ Package: "github.com/hughbliss/my_database/pkg/gen/db" + *schemaName,
		Features: []gen.Feature{
			gen.FeatureUpsert,
			gen.FeatureExecQuery,
			gen.FeatureModifier,
			gen.FeatureIntercept,
			gen.FeatureSnapshot,
			gen.FeatureVersionedMigration,
		}}, entc.Extensions()); err != nil {
		panic(err)
	}
}
